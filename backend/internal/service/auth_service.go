package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"time"

	"tbox-backend/config"
	"tbox-backend/internal/model"
	"tbox-backend/internal/repository"
	"tbox-backend/pkg/database"
	"tbox-backend/pkg/email"
	"tbox-backend/pkg/utils"
)

type AuthService struct {
	userRepo    *repository.UserRepository
	spaceRepo   *repository.SpaceRepository
	emailSender *email.EmailSender
}

func NewAuthService() *AuthService {
	cfg := config.GetConfig()
	return &AuthService{
		userRepo:    repository.NewUserRepository(),
		spaceRepo:   repository.NewSpaceRepository(),
		emailSender: email.NewEmailSender(&cfg.Email),
	}
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Code     string `json:"code" binding:"required"`
}

type LoginRequest struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RememberMe bool   `json:"remember_me"`
}

type SendCodeRequest struct {
	Email string `json:"email" binding:"required,email"`
	Type  string `json:"type" binding:"required"` // register, reset
}

type ResetPasswordRequest struct {
	Email       string `json:"email" binding:"required,email"`
	Code        string `json:"code" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

type LoginResponse struct {
	Token string      `json:"token"`
	User  *model.User `json:"user"`
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("密码长度不能少于8位")
	}
	hasLetter := regexp.MustCompile(`[a-zA-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	if !hasLetter || !hasDigit {
		return errors.New("密码必须包含字母和数字")
	}
	return nil
}

func (s *AuthService) Register(req *RegisterRequest) (*model.User, error) {
	if err := validatePassword(req.Password); err != nil {
		return nil, err
	}

	// Verify code
	ctx := context.Background()
	rdb := database.GetRedis()
	codeKey := fmt.Sprintf("verify_code:register:%s", req.Email)
	storedCode, err := rdb.Get(ctx, codeKey).Result()
	if err != nil {
		return nil, errors.New("验证码已过期或不存在")
	}
	if storedCode != req.Code {
		return nil, errors.New("验证码错误")
	}

	// Check if email exists
	existing, _ := s.userRepo.FindByEmail(req.Email)
	if existing != nil {
		return nil, errors.New("该邮箱已注册")
	}

	// Hash password
	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("密码加密失败")
	}

	user := &model.User{
		Email:        req.Email,
		PasswordHash: hash,
		Nickname:     "用户" + fmt.Sprintf("%06d", rand.Intn(1000000)),
		Status:       "active",
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, errors.New("注册失败")
	}

	// Create user space
	space := &model.UserSpace{
		UserID:        user.ID,
		TotalCapacity: 5368709120, // 5GB
		UsedCapacity:  0,
	}
	s.spaceRepo.Create(space)

	// Delete used code
	rdb.Del(ctx, codeKey)

	return user, nil
}

func (s *AuthService) Login(req *LoginRequest, clientIP string) (*LoginResponse, error) {
	ctx := context.Background()
	rdb := database.GetRedis()

	// Check login lock
	lockKey := fmt.Sprintf("login:lock:%s", req.Email)
	locked, _ := rdb.Get(ctx, lockKey).Result()
	if locked != "" {
		return nil, errors.New("登录失败次数过多，账户已锁定15分钟")
	}

	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		s.incrementLoginFailure(ctx, req.Email)
		return nil, errors.New("邮箱或密码错误")
	}

	if user.Status != "active" {
		return nil, errors.New("账户已被禁用")
	}

	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		s.incrementLoginFailure(ctx, req.Email)
		return nil, errors.New("邮箱或密码错误")
	}

	// Clear failure count
	failKey := fmt.Sprintf("login:fail:%s", req.Email)
	rdb.Del(ctx, failKey)

	// Generate token
	token, expireDuration, err := utils.GenerateToken(user.ID, user.Email, req.RememberMe)
	if err != nil {
		return nil, errors.New("生成Token失败")
	}

	// Save session to Redis
	sessionKey := fmt.Sprintf("user:session:%s", token)
	rdb.Set(ctx, sessionKey, user.ID, expireDuration)

	// Update last login info
	now := time.Now()
	user.LastLoginAt = &now
	user.LastLoginIP = clientIP
	s.userRepo.Update(user)

	return &LoginResponse{
		Token: token,
		User:  user,
	}, nil
}

func (s *AuthService) incrementLoginFailure(ctx context.Context, email string) {
	rdb := database.GetRedis()
	failKey := fmt.Sprintf("login:fail:%s", email)
	count, _ := rdb.Incr(ctx, failKey).Result()
	rdb.Expire(ctx, failKey, 15*time.Minute)

	if count >= 5 {
		lockKey := fmt.Sprintf("login:lock:%s", email)
		rdb.Set(ctx, lockKey, "locked", 15*time.Minute)
		rdb.Del(ctx, failKey)
	}
}

func (s *AuthService) SendCode(req *SendCodeRequest) error {
	ctx := context.Background()
	rdb := database.GetRedis()

	// Rate limit: 1 code per minute
	rateLimitKey := fmt.Sprintf("verify_code:rate:%s:%s", req.Type, req.Email)
	exists, _ := rdb.Exists(ctx, rateLimitKey).Result()
	if exists > 0 {
		return errors.New("发送频率过快，请稍后再试")
	}

	// For register, check if email already exists
	if req.Type == "register" {
		existing, _ := s.userRepo.FindByEmail(req.Email)
		if existing != nil {
			return errors.New("该邮箱已注册")
		}
	}

	// For reset, check if email exists
	if req.Type == "reset" {
		_, err := s.userRepo.FindByEmail(req.Email)
		if err != nil {
			return errors.New("该邮箱未注册")
		}
	}

	// Generate 6-digit code
	code := fmt.Sprintf("%06d", rand.Intn(1000000))

	// Save code to Redis with 5-minute expiration
	codeKey := fmt.Sprintf("verify_code:%s:%s", req.Type, req.Email)
	rdb.Set(ctx, codeKey, code, 5*time.Minute)

	// Set rate limit
	rdb.Set(ctx, rateLimitKey, "1", time.Minute)

	// Send email
	if err := s.emailSender.SendVerificationCode(req.Email, code); err != nil {
		// In development, log the code
		fmt.Printf("[DEV] Verification code for %s: %s\n", req.Email, code)
		// Don't return error in dev mode so testing is possible
		return nil
	}

	return nil
}

func (s *AuthService) ResetPassword(req *ResetPasswordRequest) error {
	if err := validatePassword(req.NewPassword); err != nil {
		return err
	}

	// Verify code
	ctx := context.Background()
	rdb := database.GetRedis()
	codeKey := fmt.Sprintf("verify_code:reset:%s", req.Email)
	storedCode, err := rdb.Get(ctx, codeKey).Result()
	if err != nil {
		return errors.New("验证码已过期或不存在")
	}
	if storedCode != req.Code {
		return errors.New("验证码错误")
	}

	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return errors.New("用户不存在")
	}

	hash, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return errors.New("密码加密失败")
	}

	user.PasswordHash = hash
	if err := s.userRepo.Update(user); err != nil {
		return errors.New("密码重置失败")
	}

	rdb.Del(ctx, codeKey)
	return nil
}

func (s *AuthService) Logout(token string) error {
	ctx := context.Background()
	rdb := database.GetRedis()
	sessionKey := fmt.Sprintf("user:session:%s", token)
	return rdb.Del(ctx, sessionKey).Err()
}

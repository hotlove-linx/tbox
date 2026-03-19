package service

import (
	"context"
	"fmt"
	"time"

	"admin-backend/config"
	"admin-backend/internal/model"
	"admin-backend/internal/repository"
	"admin-backend/pkg/database"
	"admin-backend/pkg/utils"
)

const (
	AdminSessionPrefix     = "admin:session:"
	AdminLoginAttemptPrefix = "admin:login_attempt:"
	AdminLoginLockPrefix   = "admin:login_lock:"
)

type AuthService struct {
	adminRepo *repository.AdminRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		adminRepo: repository.NewAdminRepository(),
	}
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Remember bool   `json:"remember"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	Admin *model.Admin `json:"admin"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

func (s *AuthService) Login(req *LoginRequest, clientIP string) (*LoginResponse, error) {
	cfg := config.GetConfig()
	rdb := database.GetRedis()
	ctx := context.Background()

	// Check if account is locked
	lockKey := AdminLoginLockPrefix + req.Email
	locked, _ := rdb.Exists(ctx, lockKey).Result()
	if locked > 0 {
		ttl, _ := rdb.TTL(ctx, lockKey).Result()
		return nil, fmt.Errorf("account is locked, please try again after %d minutes", int(ttl.Minutes())+1)
	}

	admin, err := s.adminRepo.FindByEmail(req.Email)
	if err != nil {
		s.incrementLoginAttempts(req.Email, cfg)
		return nil, fmt.Errorf("invalid email or password")
	}

	if admin.Status != "active" {
		return nil, fmt.Errorf("account is disabled")
	}

	if !utils.CheckPassword(req.Password, admin.PasswordHash) {
		s.incrementLoginAttempts(req.Email, cfg)
		return nil, fmt.Errorf("invalid email or password")
	}

	// Clear login attempts on success
	attemptKey := AdminLoginAttemptPrefix + req.Email
	rdb.Del(ctx, attemptKey)

	// Generate token
	token, expireDuration, err := utils.GenerateAdminToken(admin.ID, admin.Email, req.Remember)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token")
	}

	// Store session in Redis
	sessionKey := fmt.Sprintf("%s%d", AdminSessionPrefix, admin.ID)
	rdb.Set(ctx, sessionKey, token, expireDuration)

	// Update last login info
	now := time.Now()
	admin.LastLoginAt = &now
	admin.LastLoginIP = clientIP
	s.adminRepo.Update(admin)

	// Load roles
	admin, _ = s.adminRepo.FindByID(admin.ID)

	return &LoginResponse{
		Token: token,
		Admin: admin,
	}, nil
}

func (s *AuthService) incrementLoginAttempts(email string, cfg *config.Config) {
	rdb := database.GetRedis()
	ctx := context.Background()
	attemptKey := AdminLoginAttemptPrefix + email

	attempts, _ := rdb.Incr(ctx, attemptKey).Result()
	rdb.Expire(ctx, attemptKey, 30*time.Minute)

	if int(attempts) >= cfg.Admin.LoginMaxAttempts {
		lockKey := AdminLoginLockPrefix + email
		rdb.Set(ctx, lockKey, "1", time.Duration(cfg.Admin.LoginLockMinutes)*time.Minute)
		rdb.Del(ctx, attemptKey)
	}
}

func (s *AuthService) Logout(adminID uint64) error {
	rdb := database.GetRedis()
	ctx := context.Background()
	sessionKey := fmt.Sprintf("%s%d", AdminSessionPrefix, adminID)
	return rdb.Del(ctx, sessionKey).Err()
}

func (s *AuthService) ChangePassword(adminID uint64, req *ChangePasswordRequest) error {
	if err := utils.ValidateAdminPassword(req.NewPassword); err != nil {
		return err
	}

	admin, err := s.adminRepo.FindByID(adminID)
	if err != nil {
		return fmt.Errorf("admin not found")
	}

	if !utils.CheckPassword(req.OldPassword, admin.PasswordHash) {
		return fmt.Errorf("old password is incorrect")
	}

	hash, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return fmt.Errorf("failed to hash password")
	}

	now := time.Now()
	admin.PasswordHash = hash
	admin.PasswordChangedAt = &now
	return s.adminRepo.Update(admin)
}

func (s *AuthService) GetProfile(adminID uint64) (*model.Admin, error) {
	return s.adminRepo.FindByID(adminID)
}

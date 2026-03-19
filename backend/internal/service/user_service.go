package service

import (
	"context"
	"errors"
	"fmt"

	"tbox-backend/internal/model"
	"tbox-backend/internal/repository"
	"tbox-backend/pkg/database"
	"tbox-backend/pkg/utils"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repository.NewUserRepository(),
	}
}

type UpdateProfileRequest struct {
	Nickname string `json:"nickname" binding:"required,max=50"`
}

type UpdateAvatarRequest struct {
	AvatarURL string `json:"avatar_url" binding:"required"`
}

type ChangePasswordRequest struct {
	OldPassword     string `json:"old_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

func (s *UserService) GetProfile(userID uint64) (*model.User, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	return user, nil
}

func (s *UserService) UpdateProfile(userID uint64, req *UpdateProfileRequest) (*model.User, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	user.Nickname = req.Nickname
	if err := s.userRepo.Update(user); err != nil {
		return nil, errors.New("更新失败")
	}

	return user, nil
}

func (s *UserService) UpdateAvatar(userID uint64, req *UpdateAvatarRequest) (*model.User, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	user.AvatarURL = req.AvatarURL
	if err := s.userRepo.Update(user); err != nil {
		return nil, errors.New("更新失败")
	}

	return user, nil
}

func (s *UserService) ChangePassword(userID uint64, req *ChangePasswordRequest) error {
	if req.NewPassword != req.ConfirmPassword {
		return errors.New("两次输入的密码不一致")
	}

	if err := validatePassword(req.NewPassword); err != nil {
		return err
	}

	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return errors.New("用户不存在")
	}

	if !utils.CheckPassword(req.OldPassword, user.PasswordHash) {
		return errors.New("原密码错误")
	}

	hash, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return errors.New("密码加密失败")
	}

	user.PasswordHash = hash
	if err := s.userRepo.Update(user); err != nil {
		return errors.New("修改密码失败")
	}

	return nil
}

func (s *UserService) DeleteAccount(userID uint64, token string) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return errors.New("用户不存在")
	}

	user.Status = "deleted"
	if err := s.userRepo.Update(user); err != nil {
		return errors.New("注销失败")
	}

	// Clear session
	ctx := context.Background()
	rdb := database.GetRedis()
	sessionKey := fmt.Sprintf("user:session:%s", token)
	rdb.Del(ctx, sessionKey)

	return nil
}

package service

import (
	"fmt"
	"time"

	"admin-backend/internal/model"
	"admin-backend/internal/repository"
	"admin-backend/pkg/utils"
)

type AdminService struct {
	adminRepo *repository.AdminRepository
	roleRepo  *repository.RoleRepository
}

func NewAdminService() *AdminService {
	return &AdminService{
		adminRepo: repository.NewAdminRepository(),
		roleRepo:  repository.NewRoleRepository(),
	}
}

type CreateAdminRequest struct {
	Name     string   `json:"name" binding:"required"`
	Email    string   `json:"email" binding:"required,email"`
	Password string   `json:"password" binding:"required"`
	RoleIDs  []uint64 `json:"role_ids"`
}

type UpdateAdminRequest struct {
	Name    string   `json:"name"`
	Email   string   `json:"email"`
	RoleIDs []uint64 `json:"role_ids"`
}

type AdminListRequest struct {
	Page     int    `form:"page" binding:"required,min=1"`
	PageSize int    `form:"page_size" binding:"required,min=1,max=100"`
	Keyword  string `form:"keyword"`
}

func (s *AdminService) List(req *AdminListRequest) ([]model.Admin, int64, error) {
	return s.adminRepo.List(req.Page, req.PageSize, req.Keyword)
}

func (s *AdminService) Create(req *CreateAdminRequest) (*model.Admin, error) {
	if err := utils.ValidateAdminPassword(req.Password); err != nil {
		return nil, err
	}

	// Check if email already exists
	existing, _ := s.adminRepo.FindByEmail(req.Email)
	if existing != nil {
		return nil, fmt.Errorf("email already exists")
	}

	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password")
	}

	now := time.Now()
	admin := &model.Admin{
		Name:              req.Name,
		Email:             req.Email,
		PasswordHash:      hash,
		Status:            "active",
		PasswordChangedAt: &now,
	}

	if err := s.adminRepo.Create(admin); err != nil {
		return nil, err
	}

	if len(req.RoleIDs) > 0 {
		if err := s.adminRepo.SetRoles(admin.ID, req.RoleIDs); err != nil {
			return nil, err
		}
	}

	return s.adminRepo.FindByID(admin.ID)
}

func (s *AdminService) Update(id uint64, req *UpdateAdminRequest) error {
	admin, err := s.adminRepo.FindByID(id)
	if err != nil {
		return fmt.Errorf("admin not found")
	}

	if req.Name != "" {
		admin.Name = req.Name
	}
	if req.Email != "" {
		admin.Email = req.Email
	}

	if err := s.adminRepo.Update(admin); err != nil {
		return err
	}

	if req.RoleIDs != nil {
		return s.adminRepo.SetRoles(id, req.RoleIDs)
	}

	return nil
}

func (s *AdminService) ResetPassword(id uint64) (string, error) {
	admin, err := s.adminRepo.FindByID(id)
	if err != nil {
		return "", fmt.Errorf("admin not found")
	}

	newPassword := "Reset@" + time.Now().Format("20060102")
	hash, err := utils.HashPassword(newPassword)
	if err != nil {
		return "", fmt.Errorf("failed to hash password")
	}

	now := time.Now()
	admin.PasswordHash = hash
	admin.PasswordChangedAt = &now
	if err := s.adminRepo.Update(admin); err != nil {
		return "", err
	}

	return newPassword, nil
}

func (s *AdminService) ToggleStatus(id uint64, status string) error {
	admin, err := s.adminRepo.FindByID(id)
	if err != nil {
		return fmt.Errorf("admin not found")
	}
	admin.Status = status
	return s.adminRepo.Update(admin)
}

func (s *AdminService) Delete(id uint64) error {
	return s.adminRepo.Delete(id)
}

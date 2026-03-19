package service

import (
	"fmt"

	"admin-backend/internal/model"
	"admin-backend/internal/repository"
)

type RoleService struct {
	repo *repository.RoleRepository
}

func NewRoleService() *RoleService {
	return &RoleService{repo: repository.NewRoleRepository()}
}

type RoleRequest struct {
	Name        string `json:"name" binding:"required"`
	Code        string `json:"code" binding:"required"`
	Description string `json:"description"`
}

type SetPermissionsRequest struct {
	PermissionIDs []uint64 `json:"permission_ids" binding:"required"`
}

func (s *RoleService) List() ([]model.Role, error) {
	return s.repo.List()
}

func (s *RoleService) Create(req *RoleRequest) (*model.Role, error) {
	existing, _ := s.repo.FindByCode(req.Code)
	if existing != nil {
		return nil, fmt.Errorf("role code already exists")
	}

	role := &model.Role{
		Name:        req.Name,
		Code:        req.Code,
		Description: req.Description,
		IsSystem:    false,
	}
	err := s.repo.Create(role)
	return role, err
}

func (s *RoleService) Update(id uint64, req *RoleRequest) error {
	role, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("role not found")
	}

	role.Name = req.Name
	role.Description = req.Description
	return s.repo.Update(role)
}

func (s *RoleService) Delete(id uint64) error {
	role, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("role not found")
	}

	if role.IsSystem {
		return fmt.Errorf("cannot delete system built-in role")
	}

	hasAdmins, err := s.repo.HasAdmins(id)
	if err != nil {
		return err
	}
	if hasAdmins {
		return fmt.Errorf("cannot delete role with associated admins")
	}

	return s.repo.Delete(id)
}

func (s *RoleService) ListPermissions() ([]model.Permission, error) {
	return s.repo.ListPermissions()
}

func (s *RoleService) SetPermissions(roleID uint64, permissionIDs []uint64) error {
	_, err := s.repo.FindByID(roleID)
	if err != nil {
		return fmt.Errorf("role not found")
	}
	return s.repo.SetPermissions(roleID, permissionIDs)
}

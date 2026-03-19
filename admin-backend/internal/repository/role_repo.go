package repository

import (
	"admin-backend/internal/model"
	"admin-backend/pkg/database"

	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository() *RoleRepository {
	return &RoleRepository{db: database.GetDB()}
}

func (r *RoleRepository) FindByID(id uint64) (*model.Role, error) {
	var role model.Role
	err := r.db.Preload("Permissions").Where("id = ?", id).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepository) FindByCode(code string) (*model.Role, error) {
	var role model.Role
	err := r.db.Where("code = ?", code).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepository) List() ([]model.Role, error) {
	var roles []model.Role
	err := r.db.Preload("Permissions").Order("id ASC").Find(&roles).Error
	return roles, err
}

func (r *RoleRepository) Create(role *model.Role) error {
	return r.db.Create(role).Error
}

func (r *RoleRepository) Update(role *model.Role) error {
	return r.db.Save(role).Error
}

func (r *RoleRepository) Delete(id uint64) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("role_id = ?", id).Delete(&model.RolePermission{}).Error; err != nil {
			return err
		}
		if err := tx.Where("role_id = ?", id).Delete(&model.AdminRole{}).Error; err != nil {
			return err
		}
		return tx.Delete(&model.Role{}, id).Error
	})
}

func (r *RoleRepository) SetPermissions(roleID uint64, permissionIDs []uint64) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("role_id = ?", roleID).Delete(&model.RolePermission{}).Error; err != nil {
			return err
		}
		for _, pid := range permissionIDs {
			rp := model.RolePermission{RoleID: roleID, PermissionID: pid}
			if err := tx.Create(&rp).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *RoleRepository) ListPermissions() ([]model.Permission, error) {
	var permissions []model.Permission
	err := r.db.Order("module ASC, action ASC").Find(&permissions).Error
	return permissions, err
}

func (r *RoleRepository) HasAdmins(roleID uint64) (bool, error) {
	var count int64
	err := r.db.Model(&model.AdminRole{}).Where("role_id = ?", roleID).Count(&count).Error
	return count > 0, err
}

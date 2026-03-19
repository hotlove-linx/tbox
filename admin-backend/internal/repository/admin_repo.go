package repository

import (
	"admin-backend/internal/model"
	"admin-backend/pkg/database"

	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository() *AdminRepository {
	return &AdminRepository{db: database.GetDB()}
}

func (r *AdminRepository) FindByEmail(email string) (*model.Admin, error) {
	var admin model.Admin
	err := r.db.Where("email = ?", email).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *AdminRepository) FindByID(id uint64) (*model.Admin, error) {
	var admin model.Admin
	err := r.db.Preload("Roles").Where("id = ?", id).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *AdminRepository) Create(admin *model.Admin) error {
	return r.db.Create(admin).Error
}

func (r *AdminRepository) Update(admin *model.Admin) error {
	return r.db.Save(admin).Error
}

func (r *AdminRepository) Delete(id uint64) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("admin_id = ?", id).Delete(&model.AdminRole{}).Error; err != nil {
			return err
		}
		return tx.Delete(&model.Admin{}, id).Error
	})
}

func (r *AdminRepository) List(page, pageSize int, keyword string) ([]model.Admin, int64, error) {
	var admins []model.Admin
	var total int64

	query := r.db.Model(&model.Admin{})
	if keyword != "" {
		query = query.Where("name LIKE ? OR email LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	query.Count(&total)
	err := query.Preload("Roles").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Order("id DESC").
		Find(&admins).Error

	return admins, total, err
}

func (r *AdminRepository) SetRoles(adminID uint64, roleIDs []uint64) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("admin_id = ?", adminID).Delete(&model.AdminRole{}).Error; err != nil {
			return err
		}
		for _, roleID := range roleIDs {
			ar := model.AdminRole{AdminID: adminID, RoleID: roleID}
			if err := tx.Create(&ar).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

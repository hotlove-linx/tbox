package repository

import (
	"admin-backend/internal/model"
	"admin-backend/pkg/database"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{db: database.GetDB()}
}

func (r *CategoryRepository) List() ([]model.Category, error) {
	var categories []model.Category
	err := r.db.Order("sort_order ASC, id ASC").Find(&categories).Error
	return categories, err
}

func (r *CategoryRepository) FindByID(id uint64) (*model.Category, error) {
	var category model.Category
	err := r.db.Where("id = ?", id).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) Create(category *model.Category) error {
	return r.db.Create(category).Error
}

func (r *CategoryRepository) Update(category *model.Category) error {
	return r.db.Save(category).Error
}

func (r *CategoryRepository) Delete(id uint64) error {
	return r.db.Delete(&model.Category{}, id).Error
}

func (r *CategoryRepository) HasSoftware(id uint64) (bool, error) {
	var count int64
	err := r.db.Model(&model.Software{}).Where("category_id = ?", id).Count(&count).Error
	return count > 0, err
}

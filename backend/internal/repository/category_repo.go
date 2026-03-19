package repository

import (
	"tbox-backend/internal/model"
	"tbox-backend/pkg/database"

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
	err := r.db.Where("status = ?", "active").Order("sort_order ASC").Find(&categories).Error
	return categories, err
}

func (r *CategoryRepository) FindByID(id uint64) (*model.Category, error) {
	var category model.Category
	err := r.db.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

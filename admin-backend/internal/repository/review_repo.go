package repository

import (
	"admin-backend/internal/model"
	"admin-backend/pkg/database"

	"gorm.io/gorm"
)

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository() *ReviewRepository {
	return &ReviewRepository{db: database.GetDB()}
}

func (r *ReviewRepository) List(page, pageSize int, status string) ([]model.SoftwareReview, int64, error) {
	var list []model.SoftwareReview
	var total int64

	query := r.db.Model(&model.SoftwareReview{})
	if status != "" {
		query = query.Where("status = ?", status)
	}

	query.Count(&total)
	err := query.Preload("User").Preload("Software").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Order("id DESC").
		Find(&list).Error

	return list, total, err
}

func (r *ReviewRepository) FindByID(id uint64) (*model.SoftwareReview, error) {
	var review model.SoftwareReview
	err := r.db.Preload("User").Preload("Software").Where("id = ?", id).First(&review).Error
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *ReviewRepository) Update(review *model.SoftwareReview) error {
	return r.db.Save(review).Error
}

func (r *ReviewRepository) Delete(id uint64) error {
	return r.db.Delete(&model.SoftwareReview{}, id).Error
}

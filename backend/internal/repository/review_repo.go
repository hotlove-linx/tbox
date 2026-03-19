package repository

import (
	"tbox-backend/internal/model"
	"tbox-backend/pkg/database"

	"gorm.io/gorm"
)

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository() *ReviewRepository {
	return &ReviewRepository{db: database.GetDB()}
}

func (r *ReviewRepository) Create(review *model.SoftwareReview) error {
	return r.db.Create(review).Error
}

func (r *ReviewRepository) FindByID(id uint64) (*model.SoftwareReview, error) {
	var review model.SoftwareReview
	err := r.db.Preload("User").First(&review, id).Error
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

func (r *ReviewRepository) ListBySoftware(softwareID uint64, page, pageSize int) ([]model.SoftwareReview, int64, error) {
	var reviews []model.SoftwareReview
	var total int64

	query := r.db.Model(&model.SoftwareReview{}).Where("software_id = ? AND status = ?", softwareID, "normal")
	query.Count(&total)

	offset := (page - 1) * pageSize
	err := query.Preload("User").Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&reviews).Error
	return reviews, total, err
}

func (r *ReviewRepository) FindByUserAndSoftware(userID, softwareID uint64) (*model.SoftwareReview, error) {
	var review model.SoftwareReview
	err := r.db.Where("user_id = ? AND software_id = ?", userID, softwareID).First(&review).Error
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *ReviewRepository) GetAverageRating(softwareID uint64) (float64, error) {
	var result struct {
		AvgRating float64
	}
	err := r.db.Model(&model.SoftwareReview{}).
		Where("software_id = ? AND status = ?", softwareID, "normal").
		Select("COALESCE(AVG(rating), 0) as avg_rating").
		Scan(&result).Error
	return result.AvgRating, err
}

package repository

import (
	"admin-backend/internal/model"
	"admin-backend/pkg/database"

	"gorm.io/gorm"
)

type FeedbackRepository struct {
	db *gorm.DB
}

func NewFeedbackRepository() *FeedbackRepository {
	return &FeedbackRepository{db: database.GetDB()}
}

func (r *FeedbackRepository) List(page, pageSize int, status, feedbackType string) ([]model.Feedback, int64, error) {
	var list []model.Feedback
	var total int64

	query := r.db.Model(&model.Feedback{})
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if feedbackType != "" {
		query = query.Where("type = ?", feedbackType)
	}

	query.Count(&total)
	err := query.Preload("User").Preload("Handler").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Order("id DESC").
		Find(&list).Error

	return list, total, err
}

func (r *FeedbackRepository) FindByID(id uint64) (*model.Feedback, error) {
	var feedback model.Feedback
	err := r.db.Preload("User").Preload("Handler").Where("id = ?", id).First(&feedback).Error
	if err != nil {
		return nil, err
	}
	return &feedback, nil
}

func (r *FeedbackRepository) Update(feedback *model.Feedback) error {
	return r.db.Save(feedback).Error
}

func (r *FeedbackRepository) CountPendingReply() (int64, error) {
	var count int64
	err := r.db.Model(&model.Feedback{}).Where("status IN ?", []string{"pending", "processing"}).Count(&count).Error
	return count, err
}

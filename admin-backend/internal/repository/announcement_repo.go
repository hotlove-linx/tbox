package repository

import (
	"admin-backend/internal/model"
	"admin-backend/pkg/database"

	"gorm.io/gorm"
)

type AnnouncementRepository struct {
	db *gorm.DB
}

func NewAnnouncementRepository() *AnnouncementRepository {
	return &AnnouncementRepository{db: database.GetDB()}
}

func (r *AnnouncementRepository) List(page, pageSize int, status, announcementType string) ([]model.Announcement, int64, error) {
	var list []model.Announcement
	var total int64

	query := r.db.Model(&model.Announcement{})
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if announcementType != "" {
		query = query.Where("type = ?", announcementType)
	}

	query.Count(&total)
	err := query.Preload("Creator").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Order("id DESC").
		Find(&list).Error

	return list, total, err
}

func (r *AnnouncementRepository) FindByID(id uint64) (*model.Announcement, error) {
	var announcement model.Announcement
	err := r.db.Preload("Creator").Where("id = ?", id).First(&announcement).Error
	if err != nil {
		return nil, err
	}
	return &announcement, nil
}

func (r *AnnouncementRepository) Create(announcement *model.Announcement) error {
	return r.db.Create(announcement).Error
}

func (r *AnnouncementRepository) Update(announcement *model.Announcement) error {
	return r.db.Save(announcement).Error
}

func (r *AnnouncementRepository) Delete(id uint64) error {
	return r.db.Delete(&model.Announcement{}, id).Error
}

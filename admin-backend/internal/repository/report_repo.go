package repository

import (
	"admin-backend/internal/model"
	"admin-backend/pkg/database"

	"gorm.io/gorm"
)

type ReportRepository struct {
	db *gorm.DB
}

func NewReportRepository() *ReportRepository {
	return &ReportRepository{db: database.GetDB()}
}

func (r *ReportRepository) List(page, pageSize int, status, targetType string) ([]model.Report, int64, error) {
	var list []model.Report
	var total int64

	query := r.db.Model(&model.Report{})
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if targetType != "" {
		query = query.Where("target_type = ?", targetType)
	}

	query.Count(&total)
	err := query.Preload("Reporter").Preload("Handler").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Order("id DESC").
		Find(&list).Error

	return list, total, err
}

func (r *ReportRepository) FindByID(id uint64) (*model.Report, error) {
	var report model.Report
	err := r.db.Preload("Reporter").Preload("Handler").Where("id = ?", id).First(&report).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *ReportRepository) Update(report *model.Report) error {
	return r.db.Save(report).Error
}

func (r *ReportRepository) CountPending() (int64, error) {
	var count int64
	err := r.db.Model(&model.Report{}).Where("status = ?", "pending").Count(&count).Error
	return count, err
}

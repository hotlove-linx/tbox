package repository

import (
	"time"

	"admin-backend/internal/model"
	"admin-backend/pkg/database"

	"gorm.io/gorm"
)

type OperationLogRepository struct {
	db *gorm.DB
}

func NewOperationLogRepository() *OperationLogRepository {
	return &OperationLogRepository{db: database.GetDB()}
}

func (r *OperationLogRepository) List(page, pageSize int, adminID uint64, module, action string, startTime, endTime *time.Time) ([]model.OperationLog, int64, error) {
	var list []model.OperationLog
	var total int64

	query := r.db.Model(&model.OperationLog{})
	if adminID > 0 {
		query = query.Where("admin_id = ?", adminID)
	}
	if module != "" {
		query = query.Where("module = ?", module)
	}
	if action != "" {
		query = query.Where("action = ?", action)
	}
	if startTime != nil {
		query = query.Where("created_at >= ?", startTime)
	}
	if endTime != nil {
		query = query.Where("created_at <= ?", endTime)
	}

	query.Count(&total)
	err := query.Preload("Admin").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Order("id DESC").
		Find(&list).Error

	return list, total, err
}

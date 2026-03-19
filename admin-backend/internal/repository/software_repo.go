package repository

import (
	"time"

	"admin-backend/internal/model"
	"admin-backend/pkg/database"

	"gorm.io/gorm"
)

type SoftwareRepository struct {
	db *gorm.DB
}

func NewSoftwareRepository() *SoftwareRepository {
	return &SoftwareRepository{db: database.GetDB()}
}

func (r *SoftwareRepository) List(page, pageSize int, keyword string, categoryID uint64, auditStatus, visibility string, startTime, endTime *time.Time, sortField, sortOrder string) ([]model.Software, int64, error) {
	var list []model.Software
	var total int64

	query := r.db.Model(&model.Software{})
	if keyword != "" {
		query = query.Where("name LIKE ? OR description LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if categoryID > 0 {
		query = query.Where("category_id = ?", categoryID)
	}
	if auditStatus != "" {
		query = query.Where("audit_status = ?", auditStatus)
	}
	if visibility != "" {
		query = query.Where("visibility = ?", visibility)
	}
	if startTime != nil {
		query = query.Where("created_at >= ?", startTime)
	}
	if endTime != nil {
		query = query.Where("created_at <= ?", endTime)
	}

	query.Count(&total)

	orderClause := "id DESC"
	if sortField != "" {
		direction := "DESC"
		if sortOrder == "asc" {
			direction = "ASC"
		}
		orderClause = sortField + " " + direction
	}

	err := query.Preload("Category").Preload("Uploader").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Order(orderClause).
		Find(&list).Error

	return list, total, err
}

func (r *SoftwareRepository) FindByID(id uint64) (*model.Software, error) {
	var sw model.Software
	err := r.db.Preload("Category").Preload("Uploader").Preload("Screenshots").Preload("Tags").Preload("Versions").
		Where("id = ?", id).First(&sw).Error
	if err != nil {
		return nil, err
	}
	return &sw, nil
}

func (r *SoftwareRepository) Update(sw *model.Software) error {
	return r.db.Save(sw).Error
}

func (r *SoftwareRepository) Delete(id uint64) error {
	return r.db.Delete(&model.Software{}, id).Error
}

func (r *SoftwareRepository) CountByStatus(status string) (int64, error) {
	var count int64
	err := r.db.Model(&model.Software{}).Where("audit_status = ?", status).Count(&count).Error
	return count, err
}

func (r *SoftwareRepository) CountTotal() (int64, error) {
	var count int64
	err := r.db.Model(&model.Software{}).Count(&count).Error
	return count, err
}

func (r *SoftwareRepository) CountByCategory() ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	err := r.db.Model(&model.Software{}).
		Select("categories.name as category_name, COUNT(*) as count").
		Joins("LEFT JOIN categories ON categories.id = software.category_id").
		Group("software.category_id").
		Find(&results).Error
	return results, err
}

func (r *SoftwareRepository) TopByDownloads(limit int) ([]model.Software, error) {
	var list []model.Software
	err := r.db.Preload("Category").
		Order("download_count DESC").
		Limit(limit).
		Find(&list).Error
	return list, err
}

func (r *SoftwareRepository) BatchUpdateOnShelf(ids []uint64, onShelf bool, reason string) error {
	updates := map[string]interface{}{
		"is_on_shelf": onShelf,
	}
	if !onShelf && reason != "" {
		updates["off_shelf_reason"] = reason
	}
	return r.db.Model(&model.Software{}).Where("id IN ?", ids).Updates(updates).Error
}

func (r *SoftwareRepository) BatchDelete(ids []uint64) error {
	return r.db.Where("id IN ?", ids).Delete(&model.Software{}).Error
}

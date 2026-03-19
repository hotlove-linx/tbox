package repository

import (
	"tbox-backend/internal/model"
	"tbox-backend/pkg/database"

	"gorm.io/gorm"
)

type SoftwareRepository struct {
	db *gorm.DB
}

func NewSoftwareRepository() *SoftwareRepository {
	return &SoftwareRepository{db: database.GetDB()}
}

func (r *SoftwareRepository) Create(software *model.Software) error {
	return r.db.Create(software).Error
}

func (r *SoftwareRepository) FindByID(id uint64) (*model.Software, error) {
	var software model.Software
	err := r.db.Preload("Category").Preload("Screenshots", func(db *gorm.DB) *gorm.DB {
		return db.Order("sort_order ASC")
	}).Preload("Tags").Preload("Uploader").First(&software, id).Error
	if err != nil {
		return nil, err
	}
	return &software, nil
}

func (r *SoftwareRepository) Update(software *model.Software) error {
	return r.db.Save(software).Error
}

func (r *SoftwareRepository) Delete(id uint64) error {
	return r.db.Delete(&model.Software{}, id).Error
}

type SoftwareListParams struct {
	CategoryID uint64
	Keyword    string
	SortBy     string // hot, rating, updated, size
	PriceType  string // all, free, paid
	Page       int
	PageSize   int
}

func (r *SoftwareRepository) List(params SoftwareListParams) ([]model.Software, int64, error) {
	var software []model.Software
	var total int64

	query := r.db.Model(&model.Software{}).Where("is_on_shelf = ? AND audit_status = ? AND visibility = ?", true, "approved", "public")

	if params.CategoryID > 0 {
		query = query.Where("category_id = ?", params.CategoryID)
	}

	if params.Keyword != "" {
		query = query.Where("name LIKE ? OR description LIKE ?", "%"+params.Keyword+"%", "%"+params.Keyword+"%")
	}

	if params.PriceType == "free" {
		query = query.Where("license IN (?)", []string{"free", "open-source", "MIT", "Apache-2.0", "GPL"})
	}

	query.Count(&total)

	switch params.SortBy {
	case "hot":
		query = query.Order("download_count DESC")
	case "rating":
		query = query.Order("rating DESC")
	case "updated":
		query = query.Order("updated_at DESC")
	case "size":
		query = query.Order("size ASC")
	default:
		query = query.Order("download_count DESC")
	}

	offset := (params.Page - 1) * params.PageSize
	err := query.Preload("Category").Preload("Tags").Offset(offset).Limit(params.PageSize).Find(&software).Error
	return software, total, err
}

func (r *SoftwareRepository) Search(keyword string, page, pageSize int) ([]model.Software, int64, error) {
	var software []model.Software
	var total int64

	query := r.db.Model(&model.Software{}).
		Where("is_on_shelf = ? AND audit_status = ? AND visibility = ?", true, "approved", "public").
		Where("name LIKE ? OR description LIKE ? OR developer LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")

	query.Count(&total)

	offset := (page - 1) * pageSize
	err := query.Preload("Category").Offset(offset).Limit(pageSize).Order("download_count DESC").Find(&software).Error
	return software, total, err
}

func (r *SoftwareRepository) GetRecommended(limit int) ([]model.Software, error) {
	var software []model.Software
	err := r.db.Where("is_recommended = ? AND is_on_shelf = ? AND audit_status = ?", true, true, "approved").
		Preload("Category").Order("download_count DESC").Limit(limit).Find(&software).Error
	return software, err
}

func (r *SoftwareRepository) GetTopDownloads(limit int) ([]model.Software, error) {
	var software []model.Software
	err := r.db.Where("is_on_shelf = ? AND audit_status = ?", true, "approved").
		Order("download_count DESC").Limit(limit).Find(&software).Error
	return software, err
}

func (r *SoftwareRepository) GetByUploader(uploaderID uint64, page, pageSize int) ([]model.Software, int64, error) {
	var software []model.Software
	var total int64

	query := r.db.Model(&model.Software{}).Where("uploader_id = ?", uploaderID)
	query.Count(&total)

	offset := (page - 1) * pageSize
	err := query.Preload("Category").Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&software).Error
	return software, total, err
}

func (r *SoftwareRepository) IncrementDownloadCount(id uint64) error {
	return r.db.Model(&model.Software{}).Where("id = ?", id).UpdateColumn("download_count", gorm.Expr("download_count + ?", 1)).Error
}

func (r *SoftwareRepository) SearchSuggestions(keyword string, limit int) ([]string, error) {
	var names []string
	err := r.db.Model(&model.Software{}).
		Where("is_on_shelf = ? AND audit_status = ? AND name LIKE ?", true, "approved", "%"+keyword+"%").
		Limit(limit).Pluck("name", &names).Error
	return names, err
}

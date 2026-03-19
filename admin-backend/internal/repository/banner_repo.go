package repository

import (
	"admin-backend/internal/model"
	"admin-backend/pkg/database"

	"gorm.io/gorm"
)

type BannerRepository struct {
	db *gorm.DB
}

func NewBannerRepository() *BannerRepository {
	return &BannerRepository{db: database.GetDB()}
}

func (r *BannerRepository) List() ([]model.Banner, error) {
	var banners []model.Banner
	err := r.db.Preload("Creator").Order("sort_order ASC, id DESC").Find(&banners).Error
	return banners, err
}

func (r *BannerRepository) FindByID(id uint64) (*model.Banner, error) {
	var banner model.Banner
	err := r.db.Preload("Creator").Where("id = ?", id).First(&banner).Error
	if err != nil {
		return nil, err
	}
	return &banner, nil
}

func (r *BannerRepository) Create(banner *model.Banner) error {
	return r.db.Create(banner).Error
}

func (r *BannerRepository) Update(banner *model.Banner) error {
	return r.db.Save(banner).Error
}

func (r *BannerRepository) Delete(id uint64) error {
	return r.db.Delete(&model.Banner{}, id).Error
}

func (r *BannerRepository) BatchUpdateSort(items []struct {
	ID        uint64 `json:"id"`
	SortOrder int    `json:"sort_order"`
}) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, item := range items {
			if err := tx.Model(&model.Banner{}).Where("id = ?", item.ID).Update("sort_order", item.SortOrder).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

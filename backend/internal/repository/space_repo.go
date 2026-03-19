package repository

import (
	"tbox-backend/internal/model"
	"tbox-backend/pkg/database"

	"gorm.io/gorm"
)

type SpaceRepository struct {
	db *gorm.DB
}

func NewSpaceRepository() *SpaceRepository {
	return &SpaceRepository{db: database.GetDB()}
}

func (r *SpaceRepository) FindByUserID(userID uint64) (*model.UserSpace, error) {
	var space model.UserSpace
	err := r.db.Where("user_id = ?", userID).First(&space).Error
	if err != nil {
		return nil, err
	}
	return &space, nil
}

func (r *SpaceRepository) Create(space *model.UserSpace) error {
	return r.db.Create(space).Error
}

func (r *SpaceRepository) Update(space *model.UserSpace) error {
	return r.db.Save(space).Error
}

func (r *SpaceRepository) GetUploadStats(userID uint64) (softwareCount int64, totalDownloads int64, err error) {
	err = r.db.Model(&model.Software{}).Where("uploader_id = ?", userID).Count(&softwareCount).Error
	if err != nil {
		return
	}

	var result struct {
		TotalDownloads int64
	}
	err = r.db.Model(&model.Software{}).Where("uploader_id = ?", userID).
		Select("COALESCE(SUM(download_count), 0) as total_downloads").Scan(&result).Error
	totalDownloads = result.TotalDownloads
	return
}

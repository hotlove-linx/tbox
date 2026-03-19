package repository

import (
	"tbox-backend/internal/model"
	"tbox-backend/pkg/database"

	"gorm.io/gorm"
)

type VersionRepository struct {
	db *gorm.DB
}

func NewVersionRepository() *VersionRepository {
	return &VersionRepository{db: database.GetDB()}
}

func (r *VersionRepository) Create(version *model.SoftwareVersion) error {
	return r.db.Create(version).Error
}

func (r *VersionRepository) ListBySoftware(softwareID uint64) ([]model.SoftwareVersion, error) {
	var versions []model.SoftwareVersion
	err := r.db.Where("software_id = ?", softwareID).Order("created_at DESC").Find(&versions).Error
	return versions, err
}

func (r *VersionRepository) GetLatest(softwareID uint64) (*model.SoftwareVersion, error) {
	var version model.SoftwareVersion
	err := r.db.Where("software_id = ?", softwareID).Order("created_at DESC").First(&version).Error
	if err != nil {
		return nil, err
	}
	return &version, nil
}

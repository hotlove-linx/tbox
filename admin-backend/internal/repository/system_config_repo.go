package repository

import (
	"admin-backend/internal/model"
	"admin-backend/pkg/database"

	"gorm.io/gorm"
)

type SystemConfigRepository struct {
	db *gorm.DB
}

func NewSystemConfigRepository() *SystemConfigRepository {
	return &SystemConfigRepository{db: database.GetDB()}
}

func (r *SystemConfigRepository) ListAll() ([]model.SystemConfig, error) {
	var configs []model.SystemConfig
	err := r.db.Order("config_group ASC, id ASC").Find(&configs).Error
	return configs, err
}

func (r *SystemConfigRepository) FindByKey(key string) (*model.SystemConfig, error) {
	var cfg model.SystemConfig
	err := r.db.Where("config_key = ?", key).First(&cfg).Error
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (r *SystemConfigRepository) BatchUpdate(configs []model.SystemConfig) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, cfg := range configs {
			if err := tx.Where("config_key = ?", cfg.ConfigKey).
				Assign(model.SystemConfig{
					ConfigValue: cfg.ConfigValue,
					UpdatedBy:   cfg.UpdatedBy,
				}).FirstOrCreate(&cfg).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

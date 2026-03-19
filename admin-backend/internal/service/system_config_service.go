package service

import (
	"admin-backend/internal/model"
	"admin-backend/internal/repository"
)

type SystemConfigService struct {
	repo *repository.SystemConfigRepository
}

func NewSystemConfigService() *SystemConfigService {
	return &SystemConfigService{repo: repository.NewSystemConfigRepository()}
}

type SystemConfigUpdateRequest struct {
	Configs []struct {
		ConfigKey   string `json:"config_key" binding:"required"`
		ConfigValue string `json:"config_value"`
	} `json:"configs" binding:"required"`
}

func (s *SystemConfigService) ListAll() (map[string][]model.SystemConfig, error) {
	configs, err := s.repo.ListAll()
	if err != nil {
		return nil, err
	}

	grouped := make(map[string][]model.SystemConfig)
	for _, cfg := range configs {
		grouped[cfg.ConfigGroup] = append(grouped[cfg.ConfigGroup], cfg)
	}
	return grouped, nil
}

func (s *SystemConfigService) BatchUpdate(adminID uint64, req *SystemConfigUpdateRequest) error {
	var configs []model.SystemConfig
	for _, item := range req.Configs {
		configs = append(configs, model.SystemConfig{
			ConfigKey:   item.ConfigKey,
			ConfigValue: item.ConfigValue,
			UpdatedBy:   adminID,
		})
	}
	return s.repo.BatchUpdate(configs)
}

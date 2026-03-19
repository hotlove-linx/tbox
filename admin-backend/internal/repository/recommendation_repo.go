package repository

import (
	"admin-backend/internal/model"
	"admin-backend/pkg/database"

	"gorm.io/gorm"
)

type RecommendationRepository struct {
	db *gorm.DB
}

func NewRecommendationRepository() *RecommendationRepository {
	return &RecommendationRepository{db: database.GetDB()}
}

func (r *RecommendationRepository) List(recType string) ([]model.Recommendation, error) {
	var list []model.Recommendation
	query := r.db.Preload("Software").Preload("Creator")
	if recType != "" {
		query = query.Where("type = ?", recType)
	}
	err := query.Order("sort_order ASC, id DESC").Find(&list).Error
	return list, err
}

func (r *RecommendationRepository) FindByID(id uint64) (*model.Recommendation, error) {
	var rec model.Recommendation
	err := r.db.Preload("Software").Preload("Creator").Where("id = ?", id).First(&rec).Error
	if err != nil {
		return nil, err
	}
	return &rec, nil
}

func (r *RecommendationRepository) Create(rec *model.Recommendation) error {
	return r.db.Create(rec).Error
}

func (r *RecommendationRepository) Update(rec *model.Recommendation) error {
	return r.db.Save(rec).Error
}

func (r *RecommendationRepository) Delete(id uint64) error {
	return r.db.Delete(&model.Recommendation{}, id).Error
}

func (r *RecommendationRepository) BatchUpdateSort(items []struct {
	ID        uint64 `json:"id"`
	SortOrder int    `json:"sort_order"`
}) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, item := range items {
			if err := tx.Model(&model.Recommendation{}).Where("id = ?", item.ID).Update("sort_order", item.SortOrder).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

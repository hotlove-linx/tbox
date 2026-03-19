package repository

import (
	"admin-backend/internal/model"
	"admin-backend/pkg/database"

	"gorm.io/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository() *TagRepository {
	return &TagRepository{db: database.GetDB()}
}

func (r *TagRepository) List() ([]model.Tag, error) {
	var tags []model.Tag
	err := r.db.Order("id ASC").Find(&tags).Error
	return tags, err
}

func (r *TagRepository) FindByID(id uint64) (*model.Tag, error) {
	var tag model.Tag
	err := r.db.Where("id = ?", id).First(&tag).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *TagRepository) Create(tag *model.Tag) error {
	return r.db.Create(tag).Error
}

func (r *TagRepository) Update(tag *model.Tag) error {
	return r.db.Save(tag).Error
}

func (r *TagRepository) Delete(id uint64) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("tag_id = ?", id).Delete(&model.SoftwareTag{}).Error; err != nil {
			return err
		}
		return tx.Delete(&model.Tag{}, id).Error
	})
}

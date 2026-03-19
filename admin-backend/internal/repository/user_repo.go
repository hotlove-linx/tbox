package repository

import (
	"time"

	"admin-backend/internal/model"
	"admin-backend/pkg/database"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: database.GetDB()}
}

func (r *UserRepository) List(page, pageSize int, keyword, status string, startTime, endTime *time.Time) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	query := r.db.Model(&model.User{})
	if keyword != "" {
		query = query.Where("nickname LIKE ? OR email LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if startTime != nil {
		query = query.Where("created_at >= ?", startTime)
	}
	if endTime != nil {
		query = query.Where("created_at <= ?", endTime)
	}

	query.Count(&total)
	err := query.Offset((page - 1) * pageSize).
		Limit(pageSize).
		Order("id DESC").
		Find(&users).Error

	return users, total, err
}

func (r *UserRepository) FindByID(id uint64) (*model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) CountTotal() (int64, error) {
	var count int64
	err := r.db.Model(&model.User{}).Count(&count).Error
	return count, err
}

func (r *UserRepository) CountToday() (int64, error) {
	var count int64
	today := time.Now().Format("2006-01-02")
	err := r.db.Model(&model.User{}).Where("DATE(created_at) = ?", today).Count(&count).Error
	return count, err
}

func (r *UserRepository) CountActiveToday() (int64, error) {
	var count int64
	today := time.Now().Format("2006-01-02")
	err := r.db.Model(&model.User{}).Where("DATE(last_login_at) = ?", today).Count(&count).Error
	return count, err
}

func (r *UserRepository) CountActiveYesterday() (int64, error) {
	var count int64
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	err := r.db.Model(&model.User{}).Where("DATE(last_login_at) = ?", yesterday).Count(&count).Error
	return count, err
}

func (r *UserRepository) GetUserSoftware(userID uint64) ([]model.Software, error) {
	var list []model.Software
	err := r.db.Where("uploader_id = ?", userID).Find(&list).Error
	return list, err
}

func (r *UserRepository) GetUserReviews(userID uint64) ([]model.SoftwareReview, error) {
	var list []model.SoftwareReview
	err := r.db.Preload("Software").Where("user_id = ?", userID).Find(&list).Error
	return list, err
}

func (r *UserRepository) GetUserSpace(userID uint64) (*model.UserSpace, error) {
	var space model.UserSpace
	err := r.db.Where("user_id = ?", userID).First(&space).Error
	if err != nil {
		return nil, err
	}
	return &space, nil
}

func (r *UserRepository) UpdateUserSpace(space *model.UserSpace) error {
	return r.db.Save(space).Error
}

func (r *UserRepository) UserGrowthTrend(days int) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	err := r.db.Model(&model.User{}).
		Select("DATE(created_at) as date, COUNT(*) as count").
		Where("created_at >= DATE_SUB(NOW(), INTERVAL ? DAY)", days).
		Group("DATE(created_at)").
		Order("date ASC").
		Find(&results).Error
	return results, err
}

func (r *UserRepository) DAUTrend(days int) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	err := r.db.Model(&model.User{}).
		Select("DATE(last_login_at) as date, COUNT(*) as count").
		Where("last_login_at >= DATE_SUB(NOW(), INTERVAL ? DAY)", days).
		Group("DATE(last_login_at)").
		Order("date ASC").
		Find(&results).Error
	return results, err
}

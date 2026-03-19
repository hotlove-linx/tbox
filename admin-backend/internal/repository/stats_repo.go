package repository

import (
	"time"

	"admin-backend/internal/model"
	"admin-backend/pkg/database"

	"gorm.io/gorm"
)

type StatsRepository struct {
	db *gorm.DB
}

func NewStatsRepository() *StatsRepository {
	return &StatsRepository{db: database.GetDB()}
}

func (r *StatsRepository) TotalDownloads() (int64, error) {
	var total int64
	err := r.db.Model(&model.Software{}).Select("COALESCE(SUM(download_count), 0)").Scan(&total).Error
	return total, err
}

func (r *StatsRepository) TodayDownloads() (int64, error) {
	// This is approximate - in a real system you'd have a download_logs table
	// For now we return 0 as a placeholder
	return 0, nil
}

func (r *StatsRepository) DownloadTrend(days int) ([]map[string]interface{}, error) {
	// Placeholder: in a real system, this would query a download_logs table
	var results []map[string]interface{}
	for i := days - 1; i >= 0; i-- {
		date := time.Now().AddDate(0, 0, -i).Format("2006-01-02")
		results = append(results, map[string]interface{}{
			"date":  date,
			"count": 0,
		})
	}
	return results, nil
}

func (r *StatsRepository) SoftwareRanking(rankBy string, limit int, startTime, endTime *time.Time) ([]model.Software, error) {
	var list []model.Software
	query := r.db.Preload("Category")

	if startTime != nil {
		query = query.Where("created_at >= ?", startTime)
	}
	if endTime != nil {
		query = query.Where("created_at <= ?", endTime)
	}

	orderBy := "download_count DESC"
	switch rankBy {
	case "rating":
		orderBy = "rating DESC"
	case "newest":
		orderBy = "created_at DESC"
	case "download":
		orderBy = "download_count DESC"
	}

	err := query.Order(orderBy).Limit(limit).Find(&list).Error
	return list, err
}

package service

import (
	"admin-backend/internal/repository"
)

type DashboardService struct {
	userRepo     *repository.UserRepository
	softwareRepo *repository.SoftwareRepository
	reportRepo   *repository.ReportRepository
	feedbackRepo *repository.FeedbackRepository
	statsRepo    *repository.StatsRepository
}

func NewDashboardService() *DashboardService {
	return &DashboardService{
		userRepo:     repository.NewUserRepository(),
		softwareRepo: repository.NewSoftwareRepository(),
		reportRepo:   repository.NewReportRepository(),
		feedbackRepo: repository.NewFeedbackRepository(),
		statsRepo:    repository.NewStatsRepository(),
	}
}

type DashboardStats struct {
	TotalUsers       int64   `json:"total_users"`
	TodayNewUsers    int64   `json:"today_new_users"`
	TotalSoftware    int64   `json:"total_software"`
	PendingAudit     int64   `json:"pending_audit"`
	TodayDownloads   int64   `json:"today_downloads"`
	DownloadChange   float64 `json:"download_change"`
	TodayActiveUsers int64   `json:"today_active_users"`
	ActiveChange     float64 `json:"active_change"`
}

type TodoStats struct {
	PendingAudit    int64 `json:"pending_audit"`
	PendingReports  int64 `json:"pending_reports"`
	PendingFeedback int64 `json:"pending_feedback"`
}

func (s *DashboardService) GetStats() (*DashboardStats, error) {
	stats := &DashboardStats{}

	totalUsers, _ := s.userRepo.CountTotal()
	todayUsers, _ := s.userRepo.CountToday()
	totalSoftware, _ := s.softwareRepo.CountTotal()
	pendingAudit, _ := s.softwareRepo.CountByStatus("pending")
	todayActive, _ := s.userRepo.CountActiveToday()
	yesterdayActive, _ := s.userRepo.CountActiveYesterday()

	stats.TotalUsers = totalUsers
	stats.TodayNewUsers = todayUsers
	stats.TotalSoftware = totalSoftware
	stats.PendingAudit = pendingAudit
	stats.TodayActiveUsers = todayActive

	if yesterdayActive > 0 {
		stats.ActiveChange = float64(todayActive-yesterdayActive) / float64(yesterdayActive) * 100
	}

	return stats, nil
}

func (s *DashboardService) GetDownloadTrend(days int) ([]map[string]interface{}, error) {
	return s.statsRepo.DownloadTrend(days)
}

func (s *DashboardService) GetUserTrend(days int) ([]map[string]interface{}, error) {
	return s.userRepo.UserGrowthTrend(days)
}

func (s *DashboardService) GetCategoryDistribution() ([]map[string]interface{}, error) {
	return s.softwareRepo.CountByCategory()
}

func (s *DashboardService) GetTopSoftware(limit int) (interface{}, error) {
	return s.softwareRepo.TopByDownloads(limit)
}

func (s *DashboardService) GetTodos() (*TodoStats, error) {
	pendingAudit, _ := s.softwareRepo.CountByStatus("pending")
	pendingReports, _ := s.reportRepo.CountPending()
	pendingFeedback, _ := s.feedbackRepo.CountPendingReply()

	return &TodoStats{
		PendingAudit:    pendingAudit,
		PendingReports:  pendingReports,
		PendingFeedback: pendingFeedback,
	}, nil
}

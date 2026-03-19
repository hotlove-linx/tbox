package service

import (
	"time"

	"admin-backend/internal/model"
	"admin-backend/internal/repository"
)

type StatsService struct {
	statsRepo    *repository.StatsRepository
	userRepo     *repository.UserRepository
	softwareRepo *repository.SoftwareRepository
}

func NewStatsService() *StatsService {
	return &StatsService{
		statsRepo:    repository.NewStatsRepository(),
		userRepo:     repository.NewUserRepository(),
		softwareRepo: repository.NewSoftwareRepository(),
	}
}

type DownloadStatsRequest struct {
	Period    string `form:"period"` // day, week, month
	StartTime string `form:"start_time"`
	EndTime   string `form:"end_time"`
}

type UserStatsRequest struct {
	Period string `form:"period"` // day, week, month
	Days   int    `form:"days"`
}

type SoftwareRankingRequest struct {
	RankBy    string `form:"rank_by"` // download, rating, newest
	Limit     int    `form:"limit"`
	StartTime string `form:"start_time"`
	EndTime   string `form:"end_time"`
}

type DownloadStatsResponse struct {
	TotalDownloads int64                    `json:"total_downloads"`
	Trend          []map[string]interface{} `json:"trend"`
	CategoryDist   []map[string]interface{} `json:"category_distribution"`
}

type UserStatsResponse struct {
	TotalUsers    int64                    `json:"total_users"`
	TodayActive   int64                    `json:"today_active"`
	GrowthTrend   []map[string]interface{} `json:"growth_trend"`
	ActiveTrend   []map[string]interface{} `json:"active_trend"`
}

func (s *StatsService) GetDownloadStats(req *DownloadStatsRequest) (*DownloadStatsResponse, error) {
	days := 7
	switch req.Period {
	case "week":
		days = 7
	case "month":
		days = 30
	}

	totalDownloads, _ := s.statsRepo.TotalDownloads()
	trend, _ := s.statsRepo.DownloadTrend(days)
	categoryDist, _ := s.softwareRepo.CountByCategory()

	return &DownloadStatsResponse{
		TotalDownloads: totalDownloads,
		Trend:          trend,
		CategoryDist:   categoryDist,
	}, nil
}

func (s *StatsService) GetUserStats(req *UserStatsRequest) (*UserStatsResponse, error) {
	days := 30
	if req.Days > 0 {
		days = req.Days
	}

	totalUsers, _ := s.userRepo.CountTotal()
	todayActive, _ := s.userRepo.CountActiveToday()
	growthTrend, _ := s.userRepo.UserGrowthTrend(days)
	activeTrend, _ := s.userRepo.DAUTrend(days)

	return &UserStatsResponse{
		TotalUsers:  totalUsers,
		TodayActive: todayActive,
		GrowthTrend: growthTrend,
		ActiveTrend: activeTrend,
	}, nil
}

func (s *StatsService) GetSoftwareRanking(req *SoftwareRankingRequest) ([]model.Software, error) {
	limit := 20
	if req.Limit > 0 {
		limit = req.Limit
	}

	var startTime, endTime *time.Time
	if req.StartTime != "" {
		t, err := time.Parse("2006-01-02", req.StartTime)
		if err == nil {
			startTime = &t
		}
	}
	if req.EndTime != "" {
		t, err := time.Parse("2006-01-02", req.EndTime)
		if err == nil {
			t = t.Add(24*time.Hour - time.Second)
			endTime = &t
		}
	}

	return s.statsRepo.SoftwareRanking(req.RankBy, limit, startTime, endTime)
}

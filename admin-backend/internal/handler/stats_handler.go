package handler

import (
	"admin-backend/internal/service"
	"admin-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type StatsHandler struct {
	svc *service.StatsService
}

func NewStatsHandler() *StatsHandler {
	return &StatsHandler{svc: service.NewStatsService()}
}

func (h *StatsHandler) GetDownloadStats(c *gin.Context) {
	var req service.DownloadStatsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		// Use defaults
		req.Period = "week"
	}

	stats, err := h.svc.GetDownloadStats(&req)
	if err != nil {
		response.ServerError(c, "failed to get download stats")
		return
	}
	response.Success(c, stats)
}

func (h *StatsHandler) GetUserStats(c *gin.Context) {
	var req service.UserStatsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		req.Days = 30
	}

	stats, err := h.svc.GetUserStats(&req)
	if err != nil {
		response.ServerError(c, "failed to get user stats")
		return
	}
	response.Success(c, stats)
}

func (h *StatsHandler) GetSoftwareRanking(c *gin.Context) {
	var req service.SoftwareRankingRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		req.RankBy = "download"
		req.Limit = 20
	}

	ranking, err := h.svc.GetSoftwareRanking(&req)
	if err != nil {
		response.ServerError(c, "failed to get software ranking")
		return
	}
	response.Success(c, ranking)
}

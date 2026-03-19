package handler

import (
	"strconv"

	"admin-backend/internal/service"
	"admin-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	svc *service.DashboardService
}

func NewDashboardHandler() *DashboardHandler {
	return &DashboardHandler{svc: service.NewDashboardService()}
}

func (h *DashboardHandler) GetStats(c *gin.Context) {
	stats, err := h.svc.GetStats()
	if err != nil {
		response.ServerError(c, "failed to get stats")
		return
	}
	response.Success(c, stats)
}

func (h *DashboardHandler) GetDownloadTrend(c *gin.Context) {
	days, _ := strconv.Atoi(c.DefaultQuery("days", "7"))
	if days != 7 && days != 30 {
		days = 7
	}
	trend, err := h.svc.GetDownloadTrend(days)
	if err != nil {
		response.ServerError(c, "failed to get download trend")
		return
	}
	response.Success(c, trend)
}

func (h *DashboardHandler) GetUserTrend(c *gin.Context) {
	days, _ := strconv.Atoi(c.DefaultQuery("days", "7"))
	if days != 7 && days != 30 {
		days = 7
	}
	trend, err := h.svc.GetUserTrend(days)
	if err != nil {
		response.ServerError(c, "failed to get user trend")
		return
	}
	response.Success(c, trend)
}

func (h *DashboardHandler) GetCategoryDistribution(c *gin.Context) {
	dist, err := h.svc.GetCategoryDistribution()
	if err != nil {
		response.ServerError(c, "failed to get category distribution")
		return
	}
	response.Success(c, dist)
}

func (h *DashboardHandler) GetTopSoftware(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	top, err := h.svc.GetTopSoftware(limit)
	if err != nil {
		response.ServerError(c, "failed to get top software")
		return
	}
	response.Success(c, top)
}

func (h *DashboardHandler) GetTodos(c *gin.Context) {
	todos, err := h.svc.GetTodos()
	if err != nil {
		response.ServerError(c, "failed to get todos")
		return
	}
	response.Success(c, todos)
}

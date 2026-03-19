package handler

import (
	"strconv"

	"admin-backend/internal/middleware"
	"admin-backend/internal/service"
	"admin-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	svc *service.ReportService
}

func NewReportHandler() *ReportHandler {
	return &ReportHandler{svc: service.NewReportService()}
}

func (h *ReportHandler) List(c *gin.Context) {
	var req service.ReportListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	list, total, err := h.svc.List(&req)
	if err != nil {
		response.ServerError(c, "failed to get reports")
		return
	}
	response.SuccessPage(c, list, total, req.Page, req.PageSize)
}

func (h *ReportHandler) GetByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	report, err := h.svc.GetByID(id)
	if err != nil {
		response.NotFound(c, "report not found")
		return
	}
	response.Success(c, report)
}

func (h *ReportHandler) Resolve(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req service.ResolveReportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	adminID := middleware.GetAdminID(c)
	if err := h.svc.Resolve(id, adminID, &req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "report resolved")
}

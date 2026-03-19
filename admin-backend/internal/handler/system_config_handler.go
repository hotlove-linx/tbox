package handler

import (
	"admin-backend/internal/middleware"
	"admin-backend/internal/service"
	"admin-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type SystemConfigHandler struct {
	svc *service.SystemConfigService
}

func NewSystemConfigHandler() *SystemConfigHandler {
	return &SystemConfigHandler{svc: service.NewSystemConfigService()}
}

func (h *SystemConfigHandler) ListAll(c *gin.Context) {
	configs, err := h.svc.ListAll()
	if err != nil {
		response.ServerError(c, "failed to get system configs")
		return
	}
	response.Success(c, configs)
}

func (h *SystemConfigHandler) BatchUpdate(c *gin.Context) {
	var req service.SystemConfigUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	adminID := middleware.GetAdminID(c)
	if err := h.svc.BatchUpdate(adminID, &req); err != nil {
		response.ServerError(c, "failed to update configs")
		return
	}
	response.SuccessMessage(c, "configs updated successfully")
}

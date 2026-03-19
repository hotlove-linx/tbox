package handler

import (
	"admin-backend/internal/service"
	"admin-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type OperationLogHandler struct {
	svc *service.OperationLogService
}

func NewOperationLogHandler() *OperationLogHandler {
	return &OperationLogHandler{svc: service.NewOperationLogService()}
}

func (h *OperationLogHandler) List(c *gin.Context) {
	var req service.OperationLogListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	list, total, err := h.svc.List(&req)
	if err != nil {
		response.ServerError(c, "failed to get operation logs")
		return
	}
	response.SuccessPage(c, list, total, req.Page, req.PageSize)
}

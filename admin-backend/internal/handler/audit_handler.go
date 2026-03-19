package handler

import (
	"strconv"

	"admin-backend/internal/service"
	"admin-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuditHandler struct {
	svc *service.AuditService
}

func NewAuditHandler() *AuditHandler {
	return &AuditHandler{svc: service.NewAuditService()}
}

func (h *AuditHandler) List(c *gin.Context) {
	var req service.AuditListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	list, total, err := h.svc.List(&req)
	if err != nil {
		response.ServerError(c, "failed to get audit list")
		return
	}
	response.SuccessPage(c, list, total, req.Page, req.PageSize)
}

func (h *AuditHandler) GetByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	sw, err := h.svc.GetByID(id)
	if err != nil {
		response.NotFound(c, "software not found")
		return
	}
	response.Success(c, sw)
}

func (h *AuditHandler) Approve(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.svc.Approve(id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "software approved")
}

func (h *AuditHandler) Reject(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req service.AuditRejectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "reason is required")
		return
	}

	if err := h.svc.Reject(id, req.Reason); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "software rejected")
}

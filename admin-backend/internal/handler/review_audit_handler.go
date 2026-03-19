package handler

import (
	"strconv"

	"admin-backend/internal/service"
	"admin-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type ReviewAuditHandler struct {
	svc *service.ReviewAuditService
}

func NewReviewAuditHandler() *ReviewAuditHandler {
	return &ReviewAuditHandler{svc: service.NewReviewAuditService()}
}

func (h *ReviewAuditHandler) List(c *gin.Context) {
	var req service.ReviewAuditListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	list, total, err := h.svc.List(&req)
	if err != nil {
		response.ServerError(c, "failed to get review list")
		return
	}
	response.SuccessPage(c, list, total, req.Page, req.PageSize)
}

func (h *ReviewAuditHandler) Hide(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.svc.Hide(id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "review hidden")
}

func (h *ReviewAuditHandler) Restore(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.svc.Restore(id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "review restored")
}

func (h *ReviewAuditHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.svc.Delete(id); err != nil {
		response.ServerError(c, "failed to delete review")
		return
	}
	response.SuccessMessage(c, "review deleted")
}

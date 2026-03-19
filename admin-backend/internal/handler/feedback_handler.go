package handler

import (
	"strconv"

	"admin-backend/internal/middleware"
	"admin-backend/internal/service"
	"admin-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type FeedbackHandler struct {
	svc *service.FeedbackService
}

func NewFeedbackHandler() *FeedbackHandler {
	return &FeedbackHandler{svc: service.NewFeedbackService()}
}

func (h *FeedbackHandler) List(c *gin.Context) {
	var req service.FeedbackListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	list, total, err := h.svc.List(&req)
	if err != nil {
		response.ServerError(c, "failed to get feedbacks")
		return
	}
	response.SuccessPage(c, list, total, req.Page, req.PageSize)
}

func (h *FeedbackHandler) GetByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	feedback, err := h.svc.GetByID(id)
	if err != nil {
		response.NotFound(c, "feedback not found")
		return
	}
	response.Success(c, feedback)
}

func (h *FeedbackHandler) Assign(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req struct {
		HandlerID uint64 `json:"handler_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	if err := h.svc.Assign(id, req.HandlerID); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "feedback assigned")
}

func (h *FeedbackHandler) Reply(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req service.FeedbackReplyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	adminID := middleware.GetAdminID(c)
	if err := h.svc.Reply(id, adminID, req.Reply); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "feedback replied")
}

func (h *FeedbackHandler) Close(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.svc.Close(id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "feedback closed")
}

package handler

import (
	"strconv"

	"admin-backend/internal/middleware"
	"admin-backend/internal/service"
	"admin-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type AnnouncementHandler struct {
	svc *service.AnnouncementService
}

func NewAnnouncementHandler() *AnnouncementHandler {
	return &AnnouncementHandler{svc: service.NewAnnouncementService()}
}

func (h *AnnouncementHandler) List(c *gin.Context) {
	var req service.AnnouncementListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	list, total, err := h.svc.List(&req)
	if err != nil {
		response.ServerError(c, "failed to get announcements")
		return
	}
	response.SuccessPage(c, list, total, req.Page, req.PageSize)
}

func (h *AnnouncementHandler) Create(c *gin.Context) {
	var req service.AnnouncementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	adminID := middleware.GetAdminID(c)
	announcement, err := h.svc.Create(&req, adminID)
	if err != nil {
		response.ServerError(c, "failed to create announcement")
		return
	}
	response.Success(c, announcement)
}

func (h *AnnouncementHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req service.AnnouncementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	if err := h.svc.Update(id, &req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "announcement updated successfully")
}

func (h *AnnouncementHandler) Publish(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.svc.Publish(id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "announcement published")
}

func (h *AnnouncementHandler) Withdraw(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.svc.Withdraw(id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "announcement withdrawn")
}

func (h *AnnouncementHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.svc.Delete(id); err != nil {
		response.ServerError(c, "failed to delete announcement")
		return
	}
	response.SuccessMessage(c, "announcement deleted successfully")
}

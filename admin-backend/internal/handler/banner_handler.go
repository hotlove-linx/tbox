package handler

import (
	"strconv"

	"admin-backend/internal/middleware"
	"admin-backend/internal/service"
	"admin-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type BannerHandler struct {
	svc *service.BannerService
}

func NewBannerHandler() *BannerHandler {
	return &BannerHandler{svc: service.NewBannerService()}
}

func (h *BannerHandler) List(c *gin.Context) {
	banners, err := h.svc.List()
	if err != nil {
		response.ServerError(c, "failed to get banners")
		return
	}
	response.Success(c, banners)
}

func (h *BannerHandler) Create(c *gin.Context) {
	var req service.BannerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	adminID := middleware.GetAdminID(c)
	banner, err := h.svc.Create(&req, adminID)
	if err != nil {
		response.ServerError(c, "failed to create banner")
		return
	}
	response.Success(c, banner)
}

func (h *BannerHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req service.BannerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	if err := h.svc.Update(id, &req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "banner updated successfully")
}

func (h *BannerHandler) ToggleStatus(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	if err := h.svc.ToggleStatus(id, req.Status); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "banner status updated")
}

func (h *BannerHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.svc.Delete(id); err != nil {
		response.ServerError(c, "failed to delete banner")
		return
	}
	response.SuccessMessage(c, "banner deleted successfully")
}

func (h *BannerHandler) BatchSort(c *gin.Context) {
	var req service.BannerSortRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	if err := h.svc.BatchSort(&req); err != nil {
		response.ServerError(c, "failed to sort banners")
		return
	}
	response.SuccessMessage(c, "banners sorted successfully")
}

package handler

import (
	"strconv"

	"admin-backend/internal/service"
	"admin-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	svc *service.AdminService
}

func NewAdminHandler() *AdminHandler {
	return &AdminHandler{svc: service.NewAdminService()}
}

func (h *AdminHandler) List(c *gin.Context) {
	var req service.AdminListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	list, total, err := h.svc.List(&req)
	if err != nil {
		response.ServerError(c, "failed to get admins")
		return
	}
	response.SuccessPage(c, list, total, req.Page, req.PageSize)
}

func (h *AdminHandler) Create(c *gin.Context) {
	var req service.CreateAdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	admin, err := h.svc.Create(&req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.Success(c, admin)
}

func (h *AdminHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req service.UpdateAdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	if err := h.svc.Update(id, &req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "admin updated successfully")
}

func (h *AdminHandler) ResetPassword(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	newPassword, err := h.svc.ResetPassword(id)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.Success(c, gin.H{"new_password": newPassword})
}

func (h *AdminHandler) ToggleStatus(c *gin.Context) {
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
	response.SuccessMessage(c, "admin status updated")
}

func (h *AdminHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.svc.Delete(id); err != nil {
		response.ServerError(c, "failed to delete admin")
		return
	}
	response.SuccessMessage(c, "admin deleted successfully")
}

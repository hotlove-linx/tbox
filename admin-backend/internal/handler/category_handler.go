package handler

import (
	"strconv"

	"admin-backend/internal/service"
	"admin-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	svc *service.CategoryService
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{svc: service.NewCategoryService()}
}

func (h *CategoryHandler) List(c *gin.Context) {
	categories, err := h.svc.List()
	if err != nil {
		response.ServerError(c, "failed to get categories")
		return
	}
	response.Success(c, categories)
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var req service.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	category, err := h.svc.Create(&req)
	if err != nil {
		response.ServerError(c, "failed to create category")
		return
	}
	response.Success(c, category)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req service.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	if err := h.svc.Update(id, &req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "category updated successfully")
}

func (h *CategoryHandler) ToggleStatus(c *gin.Context) {
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
	response.SuccessMessage(c, "category status updated")
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.svc.Delete(id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "category deleted successfully")
}

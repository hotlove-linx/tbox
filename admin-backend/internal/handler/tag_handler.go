package handler

import (
	"strconv"

	"admin-backend/internal/service"
	"admin-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type TagHandler struct {
	svc *service.TagService
}

func NewTagHandler() *TagHandler {
	return &TagHandler{svc: service.NewTagService()}
}

func (h *TagHandler) List(c *gin.Context) {
	tags, err := h.svc.List()
	if err != nil {
		response.ServerError(c, "failed to get tags")
		return
	}
	response.Success(c, tags)
}

func (h *TagHandler) Create(c *gin.Context) {
	var req service.TagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	tag, err := h.svc.Create(&req)
	if err != nil {
		response.ServerError(c, "failed to create tag")
		return
	}
	response.Success(c, tag)
}

func (h *TagHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req service.TagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	if err := h.svc.Update(id, &req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "tag updated successfully")
}

func (h *TagHandler) ToggleStatus(c *gin.Context) {
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
	response.SuccessMessage(c, "tag status updated")
}

func (h *TagHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.svc.Delete(id); err != nil {
		response.ServerError(c, "failed to delete tag")
		return
	}
	response.SuccessMessage(c, "tag deleted successfully")
}

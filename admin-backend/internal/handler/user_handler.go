package handler

import (
	"strconv"

	"admin-backend/internal/service"
	"admin-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{svc: service.NewUserService()}
}

func (h *UserHandler) List(c *gin.Context) {
	var req service.UserListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	list, total, err := h.svc.List(&req)
	if err != nil {
		response.ServerError(c, "failed to get users")
		return
	}
	response.SuccessPage(c, list, total, req.Page, req.PageSize)
}

func (h *UserHandler) GetDetail(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	detail, err := h.svc.GetDetail(id)
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}
	response.Success(c, detail)
}

func (h *UserHandler) Disable(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req struct {
		Reason string `json:"reason" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "reason is required")
		return
	}

	if err := h.svc.Disable(id, req.Reason); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "user disabled")
}

func (h *UserHandler) Enable(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.svc.Enable(id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "user enabled")
}

func (h *UserHandler) AdjustSpace(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req service.AdjustSpaceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	if err := h.svc.AdjustSpace(id, req.TotalCapacity); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "user space adjusted")
}

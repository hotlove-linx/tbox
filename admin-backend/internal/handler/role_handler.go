package handler

import (
	"strconv"

	"admin-backend/internal/service"
	"admin-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	svc *service.RoleService
}

func NewRoleHandler() *RoleHandler {
	return &RoleHandler{svc: service.NewRoleService()}
}

func (h *RoleHandler) List(c *gin.Context) {
	roles, err := h.svc.List()
	if err != nil {
		response.ServerError(c, "failed to get roles")
		return
	}
	response.Success(c, roles)
}

func (h *RoleHandler) Create(c *gin.Context) {
	var req service.RoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	role, err := h.svc.Create(&req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.Success(c, role)
}

func (h *RoleHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req service.RoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	if err := h.svc.Update(id, &req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "role updated successfully")
}

func (h *RoleHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.svc.Delete(id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "role deleted successfully")
}

func (h *RoleHandler) ListPermissions(c *gin.Context) {
	permissions, err := h.svc.ListPermissions()
	if err != nil {
		response.ServerError(c, "failed to get permissions")
		return
	}
	response.Success(c, permissions)
}

func (h *RoleHandler) SetPermissions(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req service.SetPermissionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	if err := h.svc.SetPermissions(id, req.PermissionIDs); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "permissions updated successfully")
}

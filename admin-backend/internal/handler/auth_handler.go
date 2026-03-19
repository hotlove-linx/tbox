package handler

import (
	"admin-backend/internal/middleware"
	"admin-backend/internal/service"
	"admin-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	svc *service.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{svc: service.NewAuthService()}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req service.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	result, err := h.svc.Login(&req, c.ClientIP())
	if err != nil {
		response.Unauthorized(c, err.Error())
		return
	}

	response.Success(c, result)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	adminID := middleware.GetAdminID(c)
	if err := h.svc.Logout(adminID); err != nil {
		response.ServerError(c, "logout failed")
		return
	}
	response.SuccessMessage(c, "logged out successfully")
}

func (h *AuthHandler) ChangePassword(c *gin.Context) {
	var req service.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	adminID := middleware.GetAdminID(c)
	if err := h.svc.ChangePassword(adminID, &req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.SuccessMessage(c, "password changed successfully")
}

func (h *AuthHandler) GetProfile(c *gin.Context) {
	adminID := middleware.GetAdminID(c)
	admin, err := h.svc.GetProfile(adminID)
	if err != nil {
		response.NotFound(c, "admin not found")
		return
	}
	response.Success(c, admin)
}

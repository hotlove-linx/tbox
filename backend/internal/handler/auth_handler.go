package handler

import (
	"tbox-backend/internal/service"
	"tbox-backend/pkg/response"
	"tbox-backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authService: service.NewAuthService(),
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req service.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	user, err := h.authService.Register(&req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, user)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req service.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	clientIP := c.ClientIP()
	result, err := h.authService.Login(&req, clientIP)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, result)
}

func (h *AuthHandler) SendCode(c *gin.Context) {
	var req service.SendCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	if err := h.authService.SendCode(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.SuccessMessage(c, "验证码已发送")
}

func (h *AuthHandler) ResetPassword(c *gin.Context) {
	var req service.ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	if err := h.authService.ResetPassword(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.SuccessMessage(c, "密码重置成功")
}

func (h *AuthHandler) Logout(c *gin.Context) {
	token := utils.GetToken(c)
	if token == "" {
		response.BadRequest(c, "未获取到Token")
		return
	}

	if err := h.authService.Logout(token); err != nil {
		response.ServerError(c, "退出失败")
		return
	}

	response.SuccessMessage(c, "已退出登录")
}

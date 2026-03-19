package handler

import (
	"tbox-backend/internal/service"
	"tbox-backend/pkg/response"
	"tbox-backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userService: service.NewUserService(),
	}
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := utils.GetUserID(c)

	user, err := h.userService.GetProfile(userID)
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, user)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID := utils.GetUserID(c)

	var req service.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	user, err := h.userService.UpdateProfile(userID, &req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, user)
}

func (h *UserHandler) UpdateAvatar(c *gin.Context) {
	userID := utils.GetUserID(c)

	var req service.UpdateAvatarRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	user, err := h.userService.UpdateAvatar(userID, &req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, user)
}

func (h *UserHandler) ChangePassword(c *gin.Context) {
	userID := utils.GetUserID(c)

	var req service.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	if err := h.userService.ChangePassword(userID, &req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.SuccessMessage(c, "密码修改成功")
}

func (h *UserHandler) DeleteAccount(c *gin.Context) {
	userID := utils.GetUserID(c)
	token := utils.GetToken(c)

	if err := h.userService.DeleteAccount(userID, token); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.SuccessMessage(c, "账户已注销")
}

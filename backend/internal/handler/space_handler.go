package handler

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"tbox-backend/config"
	"tbox-backend/internal/service"
	"tbox-backend/pkg/response"
	"tbox-backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

type SpaceHandler struct {
	spaceService *service.SpaceService
}

func NewSpaceHandler() *SpaceHandler {
	return &SpaceHandler{
		spaceService: service.NewSpaceService(),
	}
}

func (h *SpaceHandler) GetOverview(c *gin.Context) {
	userID := utils.GetUserID(c)

	overview, err := h.spaceService.GetOverview(userID)
	if err != nil {
		response.ServerError(c, "获取空间概览失败")
		return
	}

	response.Success(c, overview)
}

func (h *SpaceHandler) GetMySoftware(c *gin.Context) {
	userID := utils.GetUserID(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	list, total, err := h.spaceService.GetMySoftware(userID, page, pageSize)
	if err != nil {
		response.ServerError(c, "获取软件列表失败")
		return
	}

	response.SuccessPage(c, list, total, page, pageSize)
}

func (h *SpaceHandler) CreateSoftware(c *gin.Context) {
	userID := utils.GetUserID(c)

	var req service.CreateSoftwareRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	software, err := h.spaceService.CreateSoftware(userID, &req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, software)
}

func (h *SpaceHandler) UpdateSoftware(c *gin.Context) {
	userID := utils.GetUserID(c)
	softwareID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的软件ID")
		return
	}

	var req service.UpdateSoftwareRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	software, err := h.spaceService.UpdateSoftware(userID, softwareID, &req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, software)
}

func (h *SpaceHandler) SubmitAudit(c *gin.Context) {
	userID := utils.GetUserID(c)
	softwareID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的软件ID")
		return
	}

	if err := h.spaceService.SubmitAudit(userID, softwareID); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.SuccessMessage(c, "已提交审核")
}

func (h *SpaceHandler) WithdrawAudit(c *gin.Context) {
	userID := utils.GetUserID(c)
	softwareID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的软件ID")
		return
	}

	if err := h.spaceService.WithdrawAudit(userID, softwareID); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.SuccessMessage(c, "已撤回审核")
}

func (h *SpaceHandler) InitUpload(c *gin.Context) {
	userID := utils.GetUserID(c)

	var req service.InitUploadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	result, err := h.spaceService.InitUpload(userID, &req)
	if err != nil {
		response.ServerError(c, err.Error())
		return
	}

	response.Success(c, result)
}

func (h *SpaceHandler) UploadChunk(c *gin.Context) {
	uploadID := c.PostForm("upload_id")
	chunkIndex := c.PostForm("chunk_index")
	if uploadID == "" || chunkIndex == "" {
		response.BadRequest(c, "缺少参数")
		return
	}

	file, err := c.FormFile("chunk")
	if err != nil {
		response.BadRequest(c, "获取文件分片失败")
		return
	}

	cfg := config.GetConfig()
	chunkDir := filepath.Join(cfg.Upload.UploadDir, "chunks", uploadID)
	if err := os.MkdirAll(chunkDir, 0755); err != nil {
		response.ServerError(c, "创建目录失败")
		return
	}

	chunkPath := filepath.Join(chunkDir, fmt.Sprintf("chunk_%s", chunkIndex))
	if err := c.SaveUploadedFile(file, chunkPath); err != nil {
		response.ServerError(c, "保存分片失败")
		return
	}

	response.SuccessMessage(c, "分片上传成功")
}

func (h *SpaceHandler) CompleteUpload(c *gin.Context) {
	userID := utils.GetUserID(c)

	var req service.CompleteUploadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	result, err := h.spaceService.CompleteUpload(userID, &req)
	if err != nil {
		response.ServerError(c, err.Error())
		return
	}

	response.Success(c, result)
}

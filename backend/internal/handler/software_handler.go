package handler

import (
	"strconv"

	"tbox-backend/internal/repository"
	"tbox-backend/internal/service"
	"tbox-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type SoftwareHandler struct {
	softwareService *service.SoftwareService
}

func NewSoftwareHandler() *SoftwareHandler {
	return &SoftwareHandler{
		softwareService: service.NewSoftwareService(),
	}
}

func (h *SoftwareHandler) GetHome(c *gin.Context) {
	data, err := h.softwareService.GetHomeData()
	if err != nil {
		response.ServerError(c, "获取首页数据失败")
		return
	}
	response.Success(c, data)
}

func (h *SoftwareHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	categoryID, _ := strconv.ParseUint(c.Query("category_id"), 10, 64)
	sortBy := c.DefaultQuery("sort_by", "hot")
	priceType := c.DefaultQuery("price_type", "all")

	params := repository.SoftwareListParams{
		CategoryID: categoryID,
		SortBy:     sortBy,
		PriceType:  priceType,
		Page:       page,
		PageSize:   pageSize,
	}

	list, total, err := h.softwareService.GetSoftwareList(params)
	if err != nil {
		response.ServerError(c, "获取软件列表失败")
		return
	}

	response.SuccessPage(c, list, total, page, pageSize)
}

func (h *SoftwareHandler) GetDetail(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的软件ID")
		return
	}

	software, err := h.softwareService.GetSoftwareDetail(id)
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, software)
}

func (h *SoftwareHandler) Search(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		response.BadRequest(c, "请输入搜索关键词")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	list, total, suggestions, err := h.softwareService.SearchSoftware(keyword, page, pageSize)
	if err != nil {
		response.ServerError(c, "搜索失败")
		return
	}

	response.Success(c, gin.H{
		"list":        list,
		"total":       total,
		"page":        page,
		"page_size":   pageSize,
		"suggestions": suggestions,
	})
}

func (h *SoftwareHandler) GetCategories(c *gin.Context) {
	categories, err := h.softwareService.GetCategories()
	if err != nil {
		response.ServerError(c, "获取分类失败")
		return
	}
	response.Success(c, categories)
}

func (h *SoftwareHandler) GetVersions(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的软件ID")
		return
	}

	versions, err := h.softwareService.GetVersions(id)
	if err != nil {
		response.ServerError(c, "获取版本历史失败")
		return
	}
	response.Success(c, versions)
}

func (h *SoftwareHandler) GetDownloadURL(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的软件ID")
		return
	}

	software, err := h.softwareService.GetDownloadURL(id)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, gin.H{
		"download_url": software.DownloadURL,
		"file_name":    software.FileName,
		"file_size":    software.Size,
		"file_hash":    software.FileHash,
	})
}

func (h *SoftwareHandler) RecordDownload(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的软件ID")
		return
	}

	if err := h.softwareService.IncrementDownloadCount(id); err != nil {
		response.ServerError(c, "记录下载失败")
		return
	}

	response.SuccessMessage(c, "已记录")
}

func (h *SoftwareHandler) CheckUpdate(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的软件ID")
		return
	}

	currentVersion := c.Query("version")
	if currentVersion == "" {
		response.BadRequest(c, "请提供当前版本号")
		return
	}

	hasUpdate, latestVersion, err := h.softwareService.CheckUpdate(id, currentVersion)
	if err != nil {
		response.ServerError(c, err.Error())
		return
	}

	response.Success(c, gin.H{
		"has_update":     hasUpdate,
		"latest_version": latestVersion,
	})
}

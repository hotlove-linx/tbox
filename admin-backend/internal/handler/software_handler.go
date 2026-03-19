package handler

import (
	"strconv"

	"admin-backend/internal/service"
	"admin-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type SoftwareHandler struct {
	svc *service.SoftwareService
}

func NewSoftwareHandler() *SoftwareHandler {
	return &SoftwareHandler{svc: service.NewSoftwareService()}
}

func (h *SoftwareHandler) List(c *gin.Context) {
	var req service.SoftwareListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	list, total, err := h.svc.List(&req)
	if err != nil {
		response.ServerError(c, "failed to get software list")
		return
	}
	response.SuccessPage(c, list, total, req.Page, req.PageSize)
}

func (h *SoftwareHandler) GetByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	sw, err := h.svc.GetByID(id)
	if err != nil {
		response.NotFound(c, "software not found")
		return
	}
	response.Success(c, sw)
}

func (h *SoftwareHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req service.SoftwareUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	if err := h.svc.Update(id, &req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "software updated successfully")
}

func (h *SoftwareHandler) OnShelf(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.svc.OnShelf(id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "software is now on shelf")
}

func (h *SoftwareHandler) OffShelf(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req struct {
		Reason string `json:"reason" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "reason is required")
		return
	}

	if err := h.svc.OffShelf(id, req.Reason); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "software is now off shelf")
}

func (h *SoftwareHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.svc.Delete(id); err != nil {
		response.ServerError(c, "failed to delete software")
		return
	}
	response.SuccessMessage(c, "software deleted successfully")
}

func (h *SoftwareHandler) BatchOperation(c *gin.Context) {
	var req service.BatchOperationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	if err := h.svc.BatchOperation(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "batch operation completed")
}

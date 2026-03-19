package handler

import (
	"strconv"

	"admin-backend/internal/middleware"
	"admin-backend/internal/service"
	"admin-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type RecommendationHandler struct {
	svc *service.RecommendationService
}

func NewRecommendationHandler() *RecommendationHandler {
	return &RecommendationHandler{svc: service.NewRecommendationService()}
}

func (h *RecommendationHandler) List(c *gin.Context) {
	recType := c.Query("type")
	list, err := h.svc.List(recType)
	if err != nil {
		response.ServerError(c, "failed to get recommendations")
		return
	}
	response.Success(c, list)
}

func (h *RecommendationHandler) Create(c *gin.Context) {
	var req service.RecommendationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	adminID := middleware.GetAdminID(c)
	rec, err := h.svc.Create(&req, adminID)
	if err != nil {
		response.ServerError(c, "failed to create recommendation")
		return
	}
	response.Success(c, rec)
}

func (h *RecommendationHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req service.RecommendationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	if err := h.svc.Update(id, &req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.SuccessMessage(c, "recommendation updated successfully")
}

func (h *RecommendationHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.svc.Delete(id); err != nil {
		response.ServerError(c, "failed to delete recommendation")
		return
	}
	response.SuccessMessage(c, "recommendation deleted successfully")
}

func (h *RecommendationHandler) BatchSort(c *gin.Context) {
	var req service.RecommendationSortRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request: "+err.Error())
		return
	}

	if err := h.svc.BatchSort(&req); err != nil {
		response.ServerError(c, "failed to sort recommendations")
		return
	}
	response.SuccessMessage(c, "recommendations sorted successfully")
}

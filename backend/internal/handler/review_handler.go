package handler

import (
	"strconv"

	"tbox-backend/internal/service"
	"tbox-backend/pkg/response"
	"tbox-backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	reviewService *service.ReviewService
}

func NewReviewHandler() *ReviewHandler {
	return &ReviewHandler{
		reviewService: service.NewReviewService(),
	}
}

func (h *ReviewHandler) GetReviews(c *gin.Context) {
	softwareID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的软件ID")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	reviews, total, err := h.reviewService.GetReviews(softwareID, page, pageSize)
	if err != nil {
		response.ServerError(c, "获取评价列表失败")
		return
	}

	response.SuccessPage(c, reviews, total, page, pageSize)
}

func (h *ReviewHandler) CreateReview(c *gin.Context) {
	userID := utils.GetUserID(c)
	softwareID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的软件ID")
		return
	}

	var req service.CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	review, err := h.reviewService.CreateReview(userID, softwareID, &req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, review)
}

func (h *ReviewHandler) UpdateReview(c *gin.Context) {
	userID := utils.GetUserID(c)
	reviewID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的评价ID")
		return
	}

	var req service.UpdateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	review, err := h.reviewService.UpdateReview(userID, reviewID, &req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, review)
}

func (h *ReviewHandler) DeleteReview(c *gin.Context) {
	userID := utils.GetUserID(c)
	reviewID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效的评价ID")
		return
	}

	if err := h.reviewService.DeleteReview(userID, reviewID); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.SuccessMessage(c, "评价已删除")
}

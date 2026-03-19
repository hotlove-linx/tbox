package service

import (
	"errors"

	"tbox-backend/internal/model"
	"tbox-backend/internal/repository"
	"tbox-backend/pkg/database"
)

type ReviewService struct {
	reviewRepo   *repository.ReviewRepository
	softwareRepo *repository.SoftwareRepository
}

func NewReviewService() *ReviewService {
	return &ReviewService{
		reviewRepo:   repository.NewReviewRepository(),
		softwareRepo: repository.NewSoftwareRepository(),
	}
}

type CreateReviewRequest struct {
	Rating  int    `json:"rating" binding:"required,min=1,max=5"`
	Content string `json:"content" binding:"required"`
}

type UpdateReviewRequest struct {
	Rating  int    `json:"rating" binding:"required,min=1,max=5"`
	Content string `json:"content" binding:"required"`
}

func (s *ReviewService) CreateReview(userID, softwareID uint64, req *CreateReviewRequest) (*model.SoftwareReview, error) {
	// Check if already reviewed
	existing, _ := s.reviewRepo.FindByUserAndSoftware(userID, softwareID)
	if existing != nil {
		return nil, errors.New("您已经评价过该软件")
	}

	review := &model.SoftwareReview{
		SoftwareID: softwareID,
		UserID:     userID,
		Rating:     req.Rating,
		Content:    req.Content,
		Status:     "normal",
	}

	if err := s.reviewRepo.Create(review); err != nil {
		return nil, errors.New("评价失败")
	}

	// Update software average rating
	s.updateSoftwareRating(softwareID)

	return review, nil
}

func (s *ReviewService) UpdateReview(userID, reviewID uint64, req *UpdateReviewRequest) (*model.SoftwareReview, error) {
	review, err := s.reviewRepo.FindByID(reviewID)
	if err != nil {
		return nil, errors.New("评价不存在")
	}

	if review.UserID != userID {
		return nil, errors.New("无权修改此评价")
	}

	review.Rating = req.Rating
	review.Content = req.Content
	if err := s.reviewRepo.Update(review); err != nil {
		return nil, errors.New("修改失败")
	}

	s.updateSoftwareRating(review.SoftwareID)

	return review, nil
}

func (s *ReviewService) DeleteReview(userID, reviewID uint64) error {
	review, err := s.reviewRepo.FindByID(reviewID)
	if err != nil {
		return errors.New("评价不存在")
	}

	if review.UserID != userID {
		return errors.New("无权删除此评价")
	}

	softwareID := review.SoftwareID
	if err := s.reviewRepo.Delete(reviewID); err != nil {
		return errors.New("删除失败")
	}

	s.updateSoftwareRating(softwareID)
	return nil
}

func (s *ReviewService) GetReviews(softwareID uint64, page, pageSize int) ([]model.SoftwareReview, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	return s.reviewRepo.ListBySoftware(softwareID, page, pageSize)
}

func (s *ReviewService) updateSoftwareRating(softwareID uint64) {
	avgRating, err := s.reviewRepo.GetAverageRating(softwareID)
	if err != nil {
		return
	}
	database.GetDB().Model(&model.Software{}).Where("id = ?", softwareID).Update("rating", avgRating)
}

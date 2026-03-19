package service

import (
	"fmt"

	"admin-backend/internal/model"
	"admin-backend/internal/repository"
)

type ReviewAuditService struct {
	repo *repository.ReviewRepository
}

func NewReviewAuditService() *ReviewAuditService {
	return &ReviewAuditService{repo: repository.NewReviewRepository()}
}

type ReviewAuditListRequest struct {
	Page     int    `form:"page" binding:"required,min=1"`
	PageSize int    `form:"page_size" binding:"required,min=1,max=100"`
	Status   string `form:"status"`
}

func (s *ReviewAuditService) List(req *ReviewAuditListRequest) ([]model.SoftwareReview, int64, error) {
	return s.repo.List(req.Page, req.PageSize, req.Status)
}

func (s *ReviewAuditService) Hide(id uint64) error {
	review, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("review not found")
	}
	review.Status = "hidden"
	return s.repo.Update(review)
}

func (s *ReviewAuditService) Restore(id uint64) error {
	review, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("review not found")
	}
	review.Status = "normal"
	return s.repo.Update(review)
}

func (s *ReviewAuditService) Delete(id uint64) error {
	return s.repo.Delete(id)
}

package service

import (
	"fmt"
	"time"

	"admin-backend/internal/model"
	"admin-backend/internal/repository"
)

type FeedbackService struct {
	repo *repository.FeedbackRepository
}

func NewFeedbackService() *FeedbackService {
	return &FeedbackService{repo: repository.NewFeedbackRepository()}
}

type FeedbackListRequest struct {
	Page     int    `form:"page" binding:"required,min=1"`
	PageSize int    `form:"page_size" binding:"required,min=1,max=100"`
	Status   string `form:"status"`
	Type     string `form:"type"`
}

type FeedbackReplyRequest struct {
	Reply string `json:"reply" binding:"required"`
}

func (s *FeedbackService) List(req *FeedbackListRequest) ([]model.Feedback, int64, error) {
	return s.repo.List(req.Page, req.PageSize, req.Status, req.Type)
}

func (s *FeedbackService) GetByID(id uint64) (*model.Feedback, error) {
	return s.repo.FindByID(id)
}

func (s *FeedbackService) Assign(id uint64, handlerID uint64) error {
	feedback, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("feedback not found")
	}

	feedback.HandlerID = &handlerID
	feedback.Status = "processing"
	return s.repo.Update(feedback)
}

func (s *FeedbackService) Reply(id uint64, adminID uint64, reply string) error {
	feedback, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("feedback not found")
	}

	now := time.Now()
	feedback.Reply = reply
	feedback.Status = "replied"
	feedback.HandlerID = &adminID
	feedback.RepliedAt = &now
	return s.repo.Update(feedback)
}

func (s *FeedbackService) Close(id uint64) error {
	feedback, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("feedback not found")
	}

	now := time.Now()
	feedback.Status = "closed"
	feedback.ClosedAt = &now
	return s.repo.Update(feedback)
}

package service

import (
	"fmt"
	"time"

	"admin-backend/internal/model"
	"admin-backend/internal/repository"
)

type AnnouncementService struct {
	repo *repository.AnnouncementRepository
}

func NewAnnouncementService() *AnnouncementService {
	return &AnnouncementService{repo: repository.NewAnnouncementRepository()}
}

type AnnouncementRequest struct {
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Type        string `json:"type" binding:"required"`
	PushMethod  string `json:"push_method"`
	TargetScope string `json:"target_scope"`
}

type AnnouncementListRequest struct {
	Page     int    `form:"page" binding:"required,min=1"`
	PageSize int    `form:"page_size" binding:"required,min=1,max=100"`
	Status   string `form:"status"`
	Type     string `form:"type"`
}

func (s *AnnouncementService) List(req *AnnouncementListRequest) ([]model.Announcement, int64, error) {
	return s.repo.List(req.Page, req.PageSize, req.Status, req.Type)
}

func (s *AnnouncementService) Create(req *AnnouncementRequest, adminID uint64) (*model.Announcement, error) {
	announcement := &model.Announcement{
		Title:       req.Title,
		Content:     req.Content,
		Type:        req.Type,
		PushMethod:  req.PushMethod,
		TargetScope: req.TargetScope,
		Status:      "draft",
		CreatedBy:   adminID,
	}
	err := s.repo.Create(announcement)
	return announcement, err
}

func (s *AnnouncementService) Update(id uint64, req *AnnouncementRequest) error {
	announcement, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("announcement not found")
	}
	announcement.Title = req.Title
	announcement.Content = req.Content
	announcement.Type = req.Type
	announcement.PushMethod = req.PushMethod
	announcement.TargetScope = req.TargetScope
	return s.repo.Update(announcement)
}

func (s *AnnouncementService) Publish(id uint64) error {
	announcement, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("announcement not found")
	}
	now := time.Now()
	announcement.Status = "published"
	announcement.PublishTime = &now
	return s.repo.Update(announcement)
}

func (s *AnnouncementService) Withdraw(id uint64) error {
	announcement, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("announcement not found")
	}
	announcement.Status = "withdrawn"
	return s.repo.Update(announcement)
}

func (s *AnnouncementService) Delete(id uint64) error {
	return s.repo.Delete(id)
}

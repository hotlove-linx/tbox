package service

import (
	"fmt"
	"time"

	"admin-backend/internal/model"
	"admin-backend/internal/repository"
)

type AuditService struct {
	softwareRepo *repository.SoftwareRepository
}

func NewAuditService() *AuditService {
	return &AuditService{softwareRepo: repository.NewSoftwareRepository()}
}

type AuditListRequest struct {
	Page        int    `form:"page" binding:"required,min=1"`
	PageSize    int    `form:"page_size" binding:"required,min=1,max=100"`
	AuditStatus string `form:"audit_status"`
	Keyword     string `form:"keyword"`
	CategoryID  uint64 `form:"category_id"`
	StartTime   string `form:"start_time"`
	EndTime     string `form:"end_time"`
}

type AuditRejectRequest struct {
	Reason string `json:"reason" binding:"required"`
}

func (s *AuditService) List(req *AuditListRequest) ([]model.Software, int64, error) {
	var startTime, endTime *time.Time
	if req.StartTime != "" {
		t, err := time.Parse("2006-01-02", req.StartTime)
		if err == nil {
			startTime = &t
		}
	}
	if req.EndTime != "" {
		t, err := time.Parse("2006-01-02", req.EndTime)
		if err == nil {
			t = t.Add(24*time.Hour - time.Second)
			endTime = &t
		}
	}

	status := req.AuditStatus
	if status == "" {
		status = "pending"
	}

	return s.softwareRepo.List(req.Page, req.PageSize, req.Keyword, req.CategoryID, status, "", startTime, endTime, "", "")
}

func (s *AuditService) GetByID(id uint64) (*model.Software, error) {
	return s.softwareRepo.FindByID(id)
}

func (s *AuditService) Approve(id uint64) error {
	sw, err := s.softwareRepo.FindByID(id)
	if err != nil {
		return fmt.Errorf("software not found")
	}

	if sw.AuditStatus != "pending" && sw.AuditStatus != "reviewing" {
		return fmt.Errorf("software is not in pending or reviewing status")
	}

	sw.AuditStatus = "approved"
	sw.IsOnShelf = true
	return s.softwareRepo.Update(sw)
}

func (s *AuditService) Reject(id uint64, reason string) error {
	sw, err := s.softwareRepo.FindByID(id)
	if err != nil {
		return fmt.Errorf("software not found")
	}

	if sw.AuditStatus != "pending" && sw.AuditStatus != "reviewing" {
		return fmt.Errorf("software is not in pending or reviewing status")
	}

	sw.AuditStatus = "rejected"
	sw.AuditRemark = reason
	return s.softwareRepo.Update(sw)
}

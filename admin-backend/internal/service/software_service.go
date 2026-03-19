package service

import (
	"fmt"
	"time"

	"admin-backend/internal/model"
	"admin-backend/internal/repository"
)

type SoftwareService struct {
	repo *repository.SoftwareRepository
}

func NewSoftwareService() *SoftwareService {
	return &SoftwareService{repo: repository.NewSoftwareRepository()}
}

type SoftwareListRequest struct {
	Page        int    `form:"page" binding:"required,min=1"`
	PageSize    int    `form:"page_size" binding:"required,min=1,max=100"`
	Keyword     string `form:"keyword"`
	CategoryID  uint64 `form:"category_id"`
	AuditStatus string `form:"audit_status"`
	Visibility  string `form:"visibility"`
	StartTime   string `form:"start_time"`
	EndTime     string `form:"end_time"`
	SortField   string `form:"sort_field"`
	SortOrder   string `form:"sort_order"`
}

type SoftwareUpdateRequest struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	Detail            string `json:"detail"`
	CategoryID        uint64 `json:"category_id"`
	Developer         string `json:"developer"`
	License           string `json:"license"`
	SystemRequirement string `json:"system_requirement"`
	Visibility        string `json:"visibility"`
}

type BatchOperationRequest struct {
	IDs       []uint64 `json:"ids" binding:"required"`
	Action    string   `json:"action" binding:"required"` // on_shelf, off_shelf, delete
	Reason    string   `json:"reason"`
}

func (s *SoftwareService) List(req *SoftwareListRequest) ([]model.Software, int64, error) {
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

	return s.repo.List(req.Page, req.PageSize, req.Keyword, req.CategoryID, req.AuditStatus, req.Visibility, startTime, endTime, req.SortField, req.SortOrder)
}

func (s *SoftwareService) GetByID(id uint64) (*model.Software, error) {
	return s.repo.FindByID(id)
}

func (s *SoftwareService) Update(id uint64, req *SoftwareUpdateRequest) error {
	sw, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("software not found")
	}

	if req.Name != "" {
		sw.Name = req.Name
	}
	if req.Description != "" {
		sw.Description = req.Description
	}
	if req.Detail != "" {
		sw.Detail = req.Detail
	}
	if req.CategoryID > 0 {
		sw.CategoryID = req.CategoryID
	}
	if req.Developer != "" {
		sw.Developer = req.Developer
	}
	if req.License != "" {
		sw.License = req.License
	}
	if req.SystemRequirement != "" {
		sw.SystemRequirement = req.SystemRequirement
	}
	if req.Visibility != "" {
		sw.Visibility = req.Visibility
	}

	return s.repo.Update(sw)
}

func (s *SoftwareService) OnShelf(id uint64) error {
	sw, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("software not found")
	}
	sw.IsOnShelf = true
	sw.OffShelfReason = ""
	return s.repo.Update(sw)
}

func (s *SoftwareService) OffShelf(id uint64, reason string) error {
	sw, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("software not found")
	}
	sw.IsOnShelf = false
	sw.OffShelfReason = reason
	return s.repo.Update(sw)
}

func (s *SoftwareService) Delete(id uint64) error {
	return s.repo.Delete(id)
}

func (s *SoftwareService) BatchOperation(req *BatchOperationRequest) error {
	switch req.Action {
	case "on_shelf":
		return s.repo.BatchUpdateOnShelf(req.IDs, true, "")
	case "off_shelf":
		return s.repo.BatchUpdateOnShelf(req.IDs, false, req.Reason)
	case "delete":
		return s.repo.BatchDelete(req.IDs)
	default:
		return fmt.Errorf("invalid action: %s", req.Action)
	}
}

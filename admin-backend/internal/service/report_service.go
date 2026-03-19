package service

import (
	"fmt"
	"time"

	"admin-backend/internal/model"
	"admin-backend/internal/repository"
)

type ReportService struct {
	repo *repository.ReportRepository
}

func NewReportService() *ReportService {
	return &ReportService{repo: repository.NewReportRepository()}
}

type ReportListRequest struct {
	Page       int    `form:"page" binding:"required,min=1"`
	PageSize   int    `form:"page_size" binding:"required,min=1,max=100"`
	Status     string `form:"status"`
	TargetType string `form:"target_type"`
}

type ResolveReportRequest struct {
	Result string `json:"result" binding:"required"` // valid or invalid
	Remark string `json:"remark"`
}

func (s *ReportService) List(req *ReportListRequest) ([]model.Report, int64, error) {
	return s.repo.List(req.Page, req.PageSize, req.Status, req.TargetType)
}

func (s *ReportService) GetByID(id uint64) (*model.Report, error) {
	return s.repo.FindByID(id)
}

func (s *ReportService) Resolve(id uint64, adminID uint64, req *ResolveReportRequest) error {
	report, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("report not found")
	}

	now := time.Now()
	report.Status = "resolved"
	report.Result = req.Result
	report.HandlerID = &adminID
	report.HandlerRemark = req.Remark
	report.ResolvedAt = &now

	return s.repo.Update(report)
}

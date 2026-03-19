package service

import (
	"time"

	"admin-backend/internal/model"
	"admin-backend/internal/repository"
)

type OperationLogService struct {
	repo *repository.OperationLogRepository
}

func NewOperationLogService() *OperationLogService {
	return &OperationLogService{repo: repository.NewOperationLogRepository()}
}

type OperationLogListRequest struct {
	Page      int    `form:"page" binding:"required,min=1"`
	PageSize  int    `form:"page_size" binding:"required,min=1,max=100"`
	AdminID   uint64 `form:"admin_id"`
	Module    string `form:"module"`
	Action    string `form:"action"`
	StartTime string `form:"start_time"`
	EndTime   string `form:"end_time"`
}

func (s *OperationLogService) List(req *OperationLogListRequest) ([]model.OperationLog, int64, error) {
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

	return s.repo.List(req.Page, req.PageSize, req.AdminID, req.Module, req.Action, startTime, endTime)
}

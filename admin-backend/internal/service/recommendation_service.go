package service

import (
	"fmt"
	"time"

	"admin-backend/internal/model"
	"admin-backend/internal/repository"
)

type RecommendationService struct {
	repo *repository.RecommendationRepository
}

func NewRecommendationService() *RecommendationService {
	return &RecommendationService{repo: repository.NewRecommendationRepository()}
}

type RecommendationRequest struct {
	SoftwareID uint64 `json:"software_id" binding:"required"`
	Type       string `json:"type" binding:"required"`
	SortOrder  int    `json:"sort_order"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
}

type RecommendationSortRequest struct {
	Items []struct {
		ID        uint64 `json:"id"`
		SortOrder int    `json:"sort_order"`
	} `json:"items" binding:"required"`
}

func (s *RecommendationService) List(recType string) ([]model.Recommendation, error) {
	return s.repo.List(recType)
}

func (s *RecommendationService) Create(req *RecommendationRequest, adminID uint64) (*model.Recommendation, error) {
	rec := &model.Recommendation{
		SoftwareID: req.SoftwareID,
		Type:       req.Type,
		SortOrder:  req.SortOrder,
		Status:     "active",
		CreatedBy:  adminID,
	}

	if req.StartTime != "" {
		t, err := time.Parse("2006-01-02 15:04:05", req.StartTime)
		if err == nil {
			rec.StartTime = &t
		}
	}
	if req.EndTime != "" {
		t, err := time.Parse("2006-01-02 15:04:05", req.EndTime)
		if err == nil {
			rec.EndTime = &t
		}
	}

	err := s.repo.Create(rec)
	return rec, err
}

func (s *RecommendationService) Update(id uint64, req *RecommendationRequest) error {
	rec, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("recommendation not found")
	}

	rec.SoftwareID = req.SoftwareID
	rec.Type = req.Type
	rec.SortOrder = req.SortOrder

	if req.StartTime != "" {
		t, err := time.Parse("2006-01-02 15:04:05", req.StartTime)
		if err == nil {
			rec.StartTime = &t
		}
	}
	if req.EndTime != "" {
		t, err := time.Parse("2006-01-02 15:04:05", req.EndTime)
		if err == nil {
			rec.EndTime = &t
		}
	}

	return s.repo.Update(rec)
}

func (s *RecommendationService) Delete(id uint64) error {
	return s.repo.Delete(id)
}

func (s *RecommendationService) BatchSort(req *RecommendationSortRequest) error {
	return s.repo.BatchUpdateSort(req.Items)
}

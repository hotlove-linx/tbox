package service

import (
	"fmt"
	"time"

	"admin-backend/internal/model"
	"admin-backend/internal/repository"
)

type BannerService struct {
	repo *repository.BannerRepository
}

func NewBannerService() *BannerService {
	return &BannerService{repo: repository.NewBannerRepository()}
}

type BannerRequest struct {
	Title      string `json:"title" binding:"required"`
	Subtitle   string `json:"subtitle"`
	ImageURL   string `json:"image_url" binding:"required"`
	LinkType   string `json:"link_type"`
	LinkTarget string `json:"link_target"`
	SortOrder  int    `json:"sort_order"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
}

type BannerSortRequest struct {
	Items []struct {
		ID        uint64 `json:"id"`
		SortOrder int    `json:"sort_order"`
	} `json:"items" binding:"required"`
}

func (s *BannerService) List() ([]model.Banner, error) {
	return s.repo.List()
}

func (s *BannerService) Create(req *BannerRequest, adminID uint64) (*model.Banner, error) {
	banner := &model.Banner{
		Title:      req.Title,
		Subtitle:   req.Subtitle,
		ImageURL:   req.ImageURL,
		LinkType:   req.LinkType,
		LinkTarget: req.LinkTarget,
		SortOrder:  req.SortOrder,
		Status:     "active",
		CreatedBy:  adminID,
	}

	if req.StartTime != "" {
		t, err := time.Parse("2006-01-02 15:04:05", req.StartTime)
		if err == nil {
			banner.StartTime = &t
		}
	}
	if req.EndTime != "" {
		t, err := time.Parse("2006-01-02 15:04:05", req.EndTime)
		if err == nil {
			banner.EndTime = &t
		}
	}

	err := s.repo.Create(banner)
	return banner, err
}

func (s *BannerService) Update(id uint64, req *BannerRequest) error {
	banner, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("banner not found")
	}

	banner.Title = req.Title
	banner.Subtitle = req.Subtitle
	banner.ImageURL = req.ImageURL
	banner.LinkType = req.LinkType
	banner.LinkTarget = req.LinkTarget
	banner.SortOrder = req.SortOrder

	if req.StartTime != "" {
		t, err := time.Parse("2006-01-02 15:04:05", req.StartTime)
		if err == nil {
			banner.StartTime = &t
		}
	}
	if req.EndTime != "" {
		t, err := time.Parse("2006-01-02 15:04:05", req.EndTime)
		if err == nil {
			banner.EndTime = &t
		}
	}

	return s.repo.Update(banner)
}

func (s *BannerService) ToggleStatus(id uint64, status string) error {
	banner, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("banner not found")
	}
	banner.Status = status
	return s.repo.Update(banner)
}

func (s *BannerService) Delete(id uint64) error {
	return s.repo.Delete(id)
}

func (s *BannerService) BatchSort(req *BannerSortRequest) error {
	return s.repo.BatchUpdateSort(req.Items)
}

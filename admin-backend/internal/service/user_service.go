package service

import (
	"fmt"
	"time"

	"admin-backend/internal/model"
	"admin-backend/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{repo: repository.NewUserRepository()}
}

type UserListRequest struct {
	Page      int    `form:"page" binding:"required,min=1"`
	PageSize  int    `form:"page_size" binding:"required,min=1,max=100"`
	Keyword   string `form:"keyword"`
	Status    string `form:"status"`
	StartTime string `form:"start_time"`
	EndTime   string `form:"end_time"`
}

type UserDetailResponse struct {
	User      *model.User             `json:"user"`
	Software  []model.Software        `json:"software"`
	Reviews   []model.SoftwareReview  `json:"reviews"`
	Space     *model.UserSpace        `json:"space"`
}

type AdjustSpaceRequest struct {
	TotalCapacity int64 `json:"total_capacity" binding:"required"`
}

func (s *UserService) List(req *UserListRequest) ([]model.User, int64, error) {
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

	return s.repo.List(req.Page, req.PageSize, req.Keyword, req.Status, startTime, endTime)
}

func (s *UserService) GetDetail(id uint64) (*UserDetailResponse, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	software, _ := s.repo.GetUserSoftware(id)
	reviews, _ := s.repo.GetUserReviews(id)
	space, _ := s.repo.GetUserSpace(id)

	return &UserDetailResponse{
		User:     user,
		Software: software,
		Reviews:  reviews,
		Space:    space,
	}, nil
}

func (s *UserService) Disable(id uint64, reason string) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	now := time.Now()
	user.Status = "disabled"
	user.DisabledReason = reason
	user.DisabledAt = &now
	return s.repo.Update(user)
}

func (s *UserService) Enable(id uint64) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	user.Status = "active"
	user.DisabledReason = ""
	user.DisabledAt = nil
	return s.repo.Update(user)
}

func (s *UserService) AdjustSpace(userID uint64, totalCapacity int64) error {
	space, err := s.repo.GetUserSpace(userID)
	if err != nil {
		// Create space record if not exists
		space = &model.UserSpace{
			UserID:        userID,
			TotalCapacity: totalCapacity,
		}
		return s.repo.UpdateUserSpace(space)
	}

	space.TotalCapacity = totalCapacity
	return s.repo.UpdateUserSpace(space)
}

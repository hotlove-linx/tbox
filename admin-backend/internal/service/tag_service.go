package service

import (
	"fmt"

	"admin-backend/internal/model"
	"admin-backend/internal/repository"
)

type TagService struct {
	repo *repository.TagRepository
}

func NewTagService() *TagService {
	return &TagService{repo: repository.NewTagRepository()}
}

type TagRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color"`
}

func (s *TagService) List() ([]model.Tag, error) {
	return s.repo.List()
}

func (s *TagService) Create(req *TagRequest) (*model.Tag, error) {
	tag := &model.Tag{
		Name:   req.Name,
		Color:  req.Color,
		Status: "active",
	}
	err := s.repo.Create(tag)
	return tag, err
}

func (s *TagService) Update(id uint64, req *TagRequest) error {
	tag, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("tag not found")
	}
	tag.Name = req.Name
	tag.Color = req.Color
	return s.repo.Update(tag)
}

func (s *TagService) ToggleStatus(id uint64, status string) error {
	tag, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("tag not found")
	}
	tag.Status = status
	return s.repo.Update(tag)
}

func (s *TagService) Delete(id uint64) error {
	return s.repo.Delete(id)
}

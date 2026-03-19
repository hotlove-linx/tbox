package service

import (
	"fmt"

	"admin-backend/internal/model"
	"admin-backend/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService() *CategoryService {
	return &CategoryService{repo: repository.NewCategoryRepository()}
}

type CategoryRequest struct {
	Name      string `json:"name" binding:"required"`
	Icon      string `json:"icon"`
	SortOrder int    `json:"sort_order"`
}

func (s *CategoryService) List() ([]model.Category, error) {
	return s.repo.List()
}

func (s *CategoryService) Create(req *CategoryRequest) (*model.Category, error) {
	category := &model.Category{
		Name:      req.Name,
		Icon:      req.Icon,
		SortOrder: req.SortOrder,
		Status:    "active",
	}
	err := s.repo.Create(category)
	return category, err
}

func (s *CategoryService) Update(id uint64, req *CategoryRequest) error {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("category not found")
	}
	category.Name = req.Name
	category.Icon = req.Icon
	category.SortOrder = req.SortOrder
	return s.repo.Update(category)
}

func (s *CategoryService) ToggleStatus(id uint64, status string) error {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("category not found")
	}
	category.Status = status
	return s.repo.Update(category)
}

func (s *CategoryService) Delete(id uint64) error {
	hasSoftware, err := s.repo.HasSoftware(id)
	if err != nil {
		return err
	}
	if hasSoftware {
		return fmt.Errorf("cannot delete category with associated software")
	}
	return s.repo.Delete(id)
}

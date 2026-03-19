package service

import (
	"errors"

	"tbox-backend/internal/model"
	"tbox-backend/internal/repository"
)

type SoftwareService struct {
	softwareRepo *repository.SoftwareRepository
	categoryRepo *repository.CategoryRepository
	versionRepo  *repository.VersionRepository
}

func NewSoftwareService() *SoftwareService {
	return &SoftwareService{
		softwareRepo: repository.NewSoftwareRepository(),
		categoryRepo: repository.NewCategoryRepository(),
		versionRepo:  repository.NewVersionRepository(),
	}
}

type HomeData struct {
	Banners    []model.Software `json:"banners"`
	Categories []model.Category `json:"categories"`
	HotList    []model.Software `json:"hot_list"`
	TopList    []model.Software `json:"top_list"`
}

func (s *SoftwareService) GetHomeData() (*HomeData, error) {
	banners, err := s.softwareRepo.GetRecommended(5)
	if err != nil {
		return nil, err
	}

	categories, err := s.categoryRepo.List()
	if err != nil {
		return nil, err
	}

	hotList, err := s.softwareRepo.GetRecommended(12)
	if err != nil {
		return nil, err
	}

	topList, err := s.softwareRepo.GetTopDownloads(10)
	if err != nil {
		return nil, err
	}

	return &HomeData{
		Banners:    banners,
		Categories: categories,
		HotList:    hotList,
		TopList:    topList,
	}, nil
}

func (s *SoftwareService) GetSoftwareList(params repository.SoftwareListParams) ([]model.Software, int64, error) {
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.PageSize <= 0 {
		params.PageSize = 20
	}
	if params.PageSize > 100 {
		params.PageSize = 100
	}
	return s.softwareRepo.List(params)
}

func (s *SoftwareService) GetSoftwareDetail(id uint64) (*model.Software, error) {
	software, err := s.softwareRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("软件不存在")
	}
	return software, nil
}

func (s *SoftwareService) SearchSoftware(keyword string, page, pageSize int) ([]model.Software, int64, []string, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}

	list, total, err := s.softwareRepo.Search(keyword, page, pageSize)
	if err != nil {
		return nil, 0, nil, err
	}

	suggestions, _ := s.softwareRepo.SearchSuggestions(keyword, 10)

	return list, total, suggestions, nil
}

func (s *SoftwareService) GetCategories() ([]model.Category, error) {
	return s.categoryRepo.List()
}

func (s *SoftwareService) GetVersions(softwareID uint64) ([]model.SoftwareVersion, error) {
	return s.versionRepo.ListBySoftware(softwareID)
}

func (s *SoftwareService) GetDownloadURL(id uint64) (*model.Software, error) {
	software, err := s.softwareRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("软件不存在")
	}
	if !software.IsOnShelf {
		return nil, errors.New("软件已下架")
	}
	return software, nil
}

func (s *SoftwareService) IncrementDownloadCount(id uint64) error {
	return s.softwareRepo.IncrementDownloadCount(id)
}

func (s *SoftwareService) CheckUpdate(id uint64, currentVersion string) (bool, *model.SoftwareVersion, error) {
	software, err := s.softwareRepo.FindByID(id)
	if err != nil {
		return false, nil, errors.New("软件不存在")
	}

	if software.Version != currentVersion {
		latestVersion, err := s.versionRepo.GetLatest(id)
		if err != nil {
			return false, nil, err
		}
		return true, latestVersion, nil
	}

	return false, nil, nil
}

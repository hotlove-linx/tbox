package service

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"tbox-backend/config"
	"tbox-backend/internal/model"
	"tbox-backend/internal/repository"

	"github.com/google/uuid"
)

type SpaceService struct {
	spaceRepo    *repository.SpaceRepository
	softwareRepo *repository.SoftwareRepository
	versionRepo  *repository.VersionRepository
}

func NewSpaceService() *SpaceService {
	return &SpaceService{
		spaceRepo:    repository.NewSpaceRepository(),
		softwareRepo: repository.NewSoftwareRepository(),
		versionRepo:  repository.NewVersionRepository(),
	}
}

type SpaceOverview struct {
	SoftwareCount  int64           `json:"software_count"`
	TotalDownloads int64           `json:"total_downloads"`
	Space          *model.UserSpace `json:"space"`
}

type CreateSoftwareRequest struct {
	Name              string `json:"name" binding:"required,max=100"`
	Description       string `json:"description"`
	Detail            string `json:"detail"`
	IconURL           string `json:"icon_url"`
	CategoryID        uint64 `json:"category_id"`
	Version           string `json:"version"`
	Size              int64  `json:"size"`
	FileName          string `json:"file_name"`
	DownloadURL       string `json:"download_url"`
	FileHash          string `json:"file_hash"`
	Developer         string `json:"developer"`
	License           string `json:"license"`
	SystemRequirement string `json:"system_requirement"`
	Visibility        string `json:"visibility"`
}

type UpdateSoftwareRequest struct {
	Name              string `json:"name" binding:"max=100"`
	Description       string `json:"description"`
	Detail            string `json:"detail"`
	IconURL           string `json:"icon_url"`
	CategoryID        uint64 `json:"category_id"`
	Version           string `json:"version"`
	License           string `json:"license"`
	SystemRequirement string `json:"system_requirement"`
	Visibility        string `json:"visibility"`
}

type InitUploadRequest struct {
	FileName   string `json:"file_name" binding:"required"`
	FileSize   int64  `json:"file_size" binding:"required"`
	ChunkSize  int64  `json:"chunk_size"`
	FileHash   string `json:"file_hash"`
}

type InitUploadResponse struct {
	UploadID   string `json:"upload_id"`
	ChunkSize  int64  `json:"chunk_size"`
	ChunkCount int    `json:"chunk_count"`
}

type CompleteUploadRequest struct {
	UploadID   string `json:"upload_id" binding:"required"`
	FileName   string `json:"file_name" binding:"required"`
	FileHash   string `json:"file_hash"`
}

type CompleteUploadResponse struct {
	FileURL  string `json:"file_url"`
	FileSize int64  `json:"file_size"`
}

func (s *SpaceService) GetOverview(userID uint64) (*SpaceOverview, error) {
	space, err := s.spaceRepo.FindByUserID(userID)
	if err != nil {
		// Create default space
		space = &model.UserSpace{
			UserID:        userID,
			TotalCapacity: 5368709120,
			UsedCapacity:  0,
		}
		s.spaceRepo.Create(space)
	}

	softwareCount, totalDownloads, err := s.spaceRepo.GetUploadStats(userID)
	if err != nil {
		return nil, err
	}

	return &SpaceOverview{
		SoftwareCount:  softwareCount,
		TotalDownloads: totalDownloads,
		Space:          space,
	}, nil
}

func (s *SpaceService) GetMySoftware(userID uint64, page, pageSize int) ([]model.Software, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	return s.softwareRepo.GetByUploader(userID, page, pageSize)
}

func (s *SpaceService) CreateSoftware(userID uint64, req *CreateSoftwareRequest) (*model.Software, error) {
	visibility := req.Visibility
	if visibility == "" {
		visibility = "public"
	}

	software := &model.Software{
		Name:              req.Name,
		Description:       req.Description,
		Detail:            req.Detail,
		IconURL:           req.IconURL,
		CategoryID:        req.CategoryID,
		Version:           req.Version,
		Size:              req.Size,
		FileName:          req.FileName,
		DownloadURL:       req.DownloadURL,
		FileHash:          req.FileHash,
		Developer:         req.Developer,
		UploaderID:        userID,
		Visibility:        visibility,
		AuditStatus:       "pending",
		IsOnShelf:         false,
		License:           req.License,
		SystemRequirement: req.SystemRequirement,
	}

	if err := s.softwareRepo.Create(software); err != nil {
		return nil, errors.New("创建软件失败")
	}

	// Create initial version
	version := &model.SoftwareVersion{
		SoftwareID:  software.ID,
		Version:     req.Version,
		Size:        req.Size,
		FileName:    req.FileName,
		DownloadURL: req.DownloadURL,
		Changelog:   "初始版本",
	}
	s.versionRepo.Create(version)

	return software, nil
}

func (s *SpaceService) UpdateSoftware(userID, softwareID uint64, req *UpdateSoftwareRequest) (*model.Software, error) {
	software, err := s.softwareRepo.FindByID(softwareID)
	if err != nil {
		return nil, errors.New("软件不存在")
	}

	if software.UploaderID != userID {
		return nil, errors.New("无权修改此软件")
	}

	if req.Name != "" {
		software.Name = req.Name
	}
	if req.Description != "" {
		software.Description = req.Description
	}
	if req.Detail != "" {
		software.Detail = req.Detail
	}
	if req.IconURL != "" {
		software.IconURL = req.IconURL
	}
	if req.CategoryID > 0 {
		software.CategoryID = req.CategoryID
	}
	if req.Version != "" {
		software.Version = req.Version
	}
	if req.License != "" {
		software.License = req.License
	}
	if req.SystemRequirement != "" {
		software.SystemRequirement = req.SystemRequirement
	}
	if req.Visibility != "" {
		software.Visibility = req.Visibility
	}

	if err := s.softwareRepo.Update(software); err != nil {
		return nil, errors.New("更新失败")
	}

	return software, nil
}

func (s *SpaceService) SubmitAudit(userID, softwareID uint64) error {
	software, err := s.softwareRepo.FindByID(softwareID)
	if err != nil {
		return errors.New("软件不存在")
	}
	if software.UploaderID != userID {
		return errors.New("无权操作此软件")
	}
	if software.AuditStatus == "reviewing" {
		return errors.New("软件正在审核中")
	}

	software.AuditStatus = "reviewing"
	return s.softwareRepo.Update(software)
}

func (s *SpaceService) WithdrawAudit(userID, softwareID uint64) error {
	software, err := s.softwareRepo.FindByID(softwareID)
	if err != nil {
		return errors.New("软件不存在")
	}
	if software.UploaderID != userID {
		return errors.New("无权操作此软件")
	}
	if software.AuditStatus != "reviewing" {
		return errors.New("当前状态无法撤回")
	}

	software.AuditStatus = "pending"
	return s.softwareRepo.Update(software)
}

func (s *SpaceService) InitUpload(userID uint64, req *InitUploadRequest) (*InitUploadResponse, error) {
	cfg := config.GetConfig()
	chunkSize := req.ChunkSize
	if chunkSize <= 0 {
		chunkSize = cfg.Upload.ChunkSize
	}

	chunkCount := int(req.FileSize / chunkSize)
	if req.FileSize%chunkSize != 0 {
		chunkCount++
	}

	uploadID := uuid.New().String()

	// Create upload directory
	uploadDir := filepath.Join(cfg.Upload.UploadDir, "chunks", uploadID)
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return nil, errors.New("创建上传目录失败")
	}

	return &InitUploadResponse{
		UploadID:   uploadID,
		ChunkSize:  chunkSize,
		ChunkCount: chunkCount,
	}, nil
}

func (s *SpaceService) CompleteUpload(userID uint64, req *CompleteUploadRequest) (*CompleteUploadResponse, error) {
	cfg := config.GetConfig()

	chunksDir := filepath.Join(cfg.Upload.UploadDir, "chunks", req.UploadID)
	finalDir := filepath.Join(cfg.Upload.UploadDir, "files")
	if err := os.MkdirAll(finalDir, 0755); err != nil {
		return nil, errors.New("创建文件目录失败")
	}

	finalPath := filepath.Join(finalDir, fmt.Sprintf("%s_%s", req.UploadID, req.FileName))
	finalFile, err := os.Create(finalPath)
	if err != nil {
		return nil, errors.New("创建目标文件失败")
	}
	defer finalFile.Close()

	// Merge chunks
	for i := 0; ; i++ {
		chunkPath := filepath.Join(chunksDir, fmt.Sprintf("chunk_%d", i))
		chunkData, err := os.ReadFile(chunkPath)
		if err != nil {
			break
		}
		finalFile.Write(chunkData)
	}

	// Get file size
	fileInfo, err := finalFile.Stat()
	if err != nil {
		return nil, errors.New("获取文件信息失败")
	}

	// Clean up chunks
	os.RemoveAll(chunksDir)

	fileURL := fmt.Sprintf("/uploads/files/%s_%s", req.UploadID, req.FileName)

	return &CompleteUploadResponse{
		FileURL:  fileURL,
		FileSize: fileInfo.Size(),
	}, nil
}

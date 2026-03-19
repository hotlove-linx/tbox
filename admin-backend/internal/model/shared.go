package model

import (
	"time"
)

// Shared models from backend - these tables already exist in the tbox database

type User struct {
	ID             uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Email          string     `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	PasswordHash   string     `gorm:"type:varchar(255);not null" json:"-"`
	Nickname       string     `gorm:"type:varchar(50)" json:"nickname"`
	AvatarURL      string     `gorm:"type:varchar(500)" json:"avatar_url"`
	Status         string     `gorm:"type:varchar(20);default:active" json:"status"`
	DisabledReason string     `gorm:"type:varchar(500)" json:"disabled_reason,omitempty"`
	DisabledAt     *time.Time `json:"disabled_at,omitempty"`
	LastLoginAt    *time.Time `json:"last_login_at,omitempty"`
	LastLoginIP    string     `gorm:"type:varchar(50)" json:"last_login_ip,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

type Software struct {
	ID                uint64               `gorm:"primaryKey;autoIncrement" json:"id"`
	Name              string               `gorm:"type:varchar(100);not null" json:"name"`
	Description       string               `gorm:"type:text" json:"description"`
	Detail            string               `gorm:"type:text" json:"detail"`
	IconURL           string               `gorm:"type:varchar(500)" json:"icon_url"`
	CategoryID        uint64               `gorm:"index" json:"category_id"`
	Version           string               `gorm:"type:varchar(20)" json:"version"`
	Size              int64                `json:"size"`
	FileName          string               `gorm:"type:varchar(255)" json:"file_name"`
	DownloadURL       string               `gorm:"type:varchar(500)" json:"download_url"`
	FileHash          string               `gorm:"type:varchar(64)" json:"file_hash"`
	Developer         string               `gorm:"type:varchar(100)" json:"developer"`
	UploaderID        uint64               `gorm:"index" json:"uploader_id"`
	Visibility        string               `gorm:"type:varchar(20);default:public" json:"visibility"`
	AuditStatus       string               `gorm:"type:varchar(20);default:pending" json:"audit_status"`
	AuditRemark       string               `gorm:"type:varchar(500)" json:"audit_remark"`
	IsRecommended     bool                 `gorm:"default:false" json:"is_recommended"`
	IsOnShelf         bool                 `gorm:"default:false" json:"is_on_shelf"`
	OffShelfReason    string               `gorm:"type:varchar(500)" json:"off_shelf_reason"`
	ScanStatus        string               `gorm:"type:varchar(20)" json:"scan_status"`
	License           string               `gorm:"type:varchar(100)" json:"license"`
	SystemRequirement string               `gorm:"type:varchar(100)" json:"system_requirement"`
	Rating            float64              `gorm:"type:decimal(2,1);default:0" json:"rating"`
	DownloadCount     int64                `gorm:"default:0" json:"download_count"`
	CreatedAt         time.Time            `json:"created_at"`
	UpdatedAt         time.Time            `json:"updated_at"`
	Category          *Category            `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Uploader          *User                `gorm:"foreignKey:UploaderID" json:"uploader,omitempty"`
	Screenshots       []SoftwareScreenshot `gorm:"foreignKey:SoftwareID" json:"screenshots,omitempty"`
	Tags              []Tag                `gorm:"many2many:software_tags;" json:"tags,omitempty"`
	Versions          []SoftwareVersion    `gorm:"foreignKey:SoftwareID" json:"versions,omitempty"`
}

func (Software) TableName() string {
	return "software"
}

type Category struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(50);not null" json:"name"`
	Icon      string    `gorm:"type:varchar(100)" json:"icon"`
	SortOrder int       `gorm:"default:0" json:"sort_order"`
	Status    string    `gorm:"type:varchar(20);default:active" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Category) TableName() string {
	return "categories"
}

type Tag struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(50);not null" json:"name"`
	Color     string    `gorm:"type:varchar(20)" json:"color"`
	Type      string    `gorm:"type:varchar(20);default:manual" json:"type"`
	Status    string    `gorm:"type:varchar(20);default:active" json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func (Tag) TableName() string {
	return "tags"
}

type SoftwareTag struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	SoftwareID uint64 `gorm:"index" json:"software_id"`
	TagID      uint64 `gorm:"index" json:"tag_id"`
}

func (SoftwareTag) TableName() string {
	return "software_tags"
}

type SoftwareScreenshot struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	SoftwareID uint64    `gorm:"index" json:"software_id"`
	ImageURL   string    `gorm:"type:varchar(500)" json:"image_url"`
	SortOrder  int       `gorm:"default:0" json:"sort_order"`
	CreatedAt  time.Time `json:"created_at"`
}

func (SoftwareScreenshot) TableName() string {
	return "software_screenshots"
}

type SoftwareVersion struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	SoftwareID  uint64    `gorm:"index" json:"software_id"`
	Version     string    `gorm:"type:varchar(20)" json:"version"`
	Size        int64     `json:"size"`
	FileName    string    `gorm:"type:varchar(255)" json:"file_name"`
	DownloadURL string    `gorm:"type:varchar(500)" json:"download_url"`
	Changelog   string    `gorm:"type:text" json:"changelog"`
	CreatedAt   time.Time `json:"created_at"`
}

func (SoftwareVersion) TableName() string {
	return "software_versions"
}

type SoftwareReview struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	SoftwareID uint64    `gorm:"index" json:"software_id"`
	UserID     uint64    `gorm:"index" json:"user_id"`
	Rating     int       `gorm:"not null" json:"rating"`
	Content    string    `gorm:"type:text" json:"content"`
	Status     string    `gorm:"type:varchar(20);default:normal" json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	User       *User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Software   *Software `gorm:"foreignKey:SoftwareID" json:"software,omitempty"`
}

func (SoftwareReview) TableName() string {
	return "software_reviews"
}

type UserSpace struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID        uint64    `gorm:"uniqueIndex" json:"user_id"`
	TotalCapacity int64     `gorm:"default:5368709120" json:"total_capacity"`
	UsedCapacity  int64     `gorm:"default:0" json:"used_capacity"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (UserSpace) TableName() string {
	return "user_spaces"
}

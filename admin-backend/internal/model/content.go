package model

import (
	"time"
)

type Banner struct {
	ID         uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Title      string     `gorm:"type:varchar(100);not null" json:"title"`
	Subtitle   string     `gorm:"type:varchar(200)" json:"subtitle"`
	ImageURL   string     `gorm:"type:varchar(500)" json:"image_url"`
	LinkType   string     `gorm:"type:varchar(20)" json:"link_type"`
	LinkTarget string     `gorm:"type:varchar(500)" json:"link_target"`
	SortOrder  int        `gorm:"default:0" json:"sort_order"`
	Status     string     `gorm:"type:varchar(20);default:active" json:"status"`
	StartTime  *time.Time `json:"start_time,omitempty"`
	EndTime    *time.Time `json:"end_time,omitempty"`
	CreatedBy  uint64     `json:"created_by"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	Creator    *Admin     `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
}

func (Banner) TableName() string {
	return "banners"
}

type Announcement struct {
	ID          uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string     `gorm:"type:varchar(200);not null" json:"title"`
	Content     string     `gorm:"type:text" json:"content"`
	Type        string     `gorm:"type:varchar(20)" json:"type"`
	PushMethod  string     `gorm:"type:varchar(20)" json:"push_method"`
	TargetScope string     `gorm:"type:varchar(20);default:all" json:"target_scope"`
	Status      string     `gorm:"type:varchar(20);default:draft" json:"status"`
	PublishTime *time.Time `json:"publish_time,omitempty"`
	CreatedBy   uint64     `json:"created_by"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Creator     *Admin     `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
}

func (Announcement) TableName() string {
	return "announcements"
}

type Recommendation struct {
	ID         uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	SoftwareID uint64     `gorm:"index" json:"software_id"`
	Type       string     `gorm:"type:varchar(20)" json:"type"`
	SortOrder  int        `gorm:"default:0" json:"sort_order"`
	StartTime  *time.Time `json:"start_time,omitempty"`
	EndTime    *time.Time `json:"end_time,omitempty"`
	Status     string     `gorm:"type:varchar(20);default:active" json:"status"`
	CreatedBy  uint64     `json:"created_by"`
	CreatedAt  time.Time  `json:"created_at"`
	Software   *Software  `gorm:"foreignKey:SoftwareID" json:"software,omitempty"`
	Creator    *Admin     `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
}

func (Recommendation) TableName() string {
	return "recommendations"
}

type SystemConfig struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	ConfigGroup string    `gorm:"type:varchar(50)" json:"config_group"`
	ConfigKey   string    `gorm:"type:varchar(100);uniqueIndex" json:"config_key"`
	ConfigValue string    `gorm:"type:text" json:"config_value"`
	Description string    `gorm:"type:varchar(200)" json:"description"`
	UpdatedBy   uint64    `json:"updated_by"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (SystemConfig) TableName() string {
	return "system_configs"
}

package model

import (
	"time"
)

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

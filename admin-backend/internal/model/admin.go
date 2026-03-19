package model

import (
	"time"
)

type Admin struct {
	ID                uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name              string     `gorm:"type:varchar(50);not null" json:"name"`
	Email             string     `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	PasswordHash      string     `gorm:"type:varchar(255);not null" json:"-"`
	AvatarURL         string     `gorm:"type:varchar(500)" json:"avatar_url"`
	Status            string     `gorm:"type:varchar(20);default:active" json:"status"`
	LastLoginAt       *time.Time `json:"last_login_at,omitempty"`
	LastLoginIP       string     `gorm:"type:varchar(50)" json:"last_login_ip,omitempty"`
	PasswordChangedAt *time.Time `json:"password_changed_at,omitempty"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	Roles             []Role     `gorm:"many2many:admin_roles;" json:"roles,omitempty"`
}

func (Admin) TableName() string {
	return "admins"
}

type Role struct {
	ID          uint64       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string       `gorm:"type:varchar(50);not null" json:"name"`
	Code        string       `gorm:"type:varchar(50);uniqueIndex;not null" json:"code"`
	Description string       `gorm:"type:varchar(200)" json:"description"`
	IsSystem    bool         `gorm:"default:false" json:"is_system"`
	CreatedAt   time.Time    `json:"created_at"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions,omitempty"`
}

func (Role) TableName() string {
	return "roles"
}

type AdminRole struct {
	ID      uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	AdminID uint64 `gorm:"index" json:"admin_id"`
	RoleID  uint64 `gorm:"index" json:"role_id"`
}

func (AdminRole) TableName() string {
	return "admin_roles"
}

type Permission struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Module      string `gorm:"type:varchar(50);not null" json:"module"`
	Action      string `gorm:"type:varchar(50);not null" json:"action"`
	Code        string `gorm:"type:varchar(100);uniqueIndex;not null" json:"code"`
	Description string `gorm:"type:varchar(200)" json:"description"`
}

func (Permission) TableName() string {
	return "permissions"
}

type RolePermission struct {
	ID           uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	RoleID       uint64 `gorm:"index" json:"role_id"`
	PermissionID uint64 `gorm:"index" json:"permission_id"`
}

func (RolePermission) TableName() string {
	return "role_permissions"
}

type OperationLog struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	AdminID    uint64    `gorm:"index" json:"admin_id"`
	Module     string    `gorm:"type:varchar(50)" json:"module"`
	Action     string    `gorm:"type:varchar(50)" json:"action"`
	TargetType string    `gorm:"type:varchar(50)" json:"target_type"`
	TargetID   uint64    `json:"target_id"`
	TargetName string    `gorm:"type:varchar(200)" json:"target_name"`
	Detail     string    `gorm:"type:text" json:"detail"`
	IP         string    `gorm:"type:varchar(50)" json:"ip"`
	CreatedAt  time.Time `json:"created_at"`
	Admin      *Admin    `gorm:"foreignKey:AdminID" json:"admin,omitempty"`
}

func (OperationLog) TableName() string {
	return "operation_logs"
}

package model

import (
	"time"
)

type Report struct {
	ID            uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	ReporterID    uint64     `gorm:"index" json:"reporter_id"`
	TargetType    string     `gorm:"type:varchar(20)" json:"target_type"`
	TargetID      uint64     `json:"target_id"`
	Reason        string     `gorm:"type:varchar(50)" json:"reason"`
	Description   string     `gorm:"type:text" json:"description"`
	Status        string     `gorm:"type:varchar(20);default:pending" json:"status"`
	Result        string     `gorm:"type:varchar(20)" json:"result"`
	HandlerID     *uint64    `json:"handler_id,omitempty"`
	HandlerRemark string     `gorm:"type:text" json:"handler_remark"`
	CreatedAt     time.Time  `json:"created_at"`
	ResolvedAt    *time.Time `json:"resolved_at,omitempty"`
	Reporter      *User      `gorm:"foreignKey:ReporterID" json:"reporter,omitempty"`
	Handler       *Admin     `gorm:"foreignKey:HandlerID" json:"handler,omitempty"`
}

func (Report) TableName() string {
	return "reports"
}

type Feedback struct {
	ID          uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uint64     `gorm:"index" json:"user_id"`
	Type        string     `gorm:"type:varchar(20)" json:"type"`
	Title       string     `gorm:"type:varchar(200)" json:"title"`
	Content     string     `gorm:"type:text" json:"content"`
	Screenshots string     `gorm:"type:text" json:"screenshots"`
	Status      string     `gorm:"type:varchar(20);default:pending" json:"status"`
	HandlerID   *uint64    `json:"handler_id,omitempty"`
	Reply       string     `gorm:"type:text" json:"reply"`
	CreatedAt   time.Time  `json:"created_at"`
	RepliedAt   *time.Time `json:"replied_at,omitempty"`
	ClosedAt    *time.Time `json:"closed_at,omitempty"`
	User        *User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Handler     *Admin     `gorm:"foreignKey:HandlerID" json:"handler,omitempty"`
}

func (Feedback) TableName() string {
	return "feedbacks"
}

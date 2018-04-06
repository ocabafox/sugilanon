package models

import (
	"time"
)

type AppUserRole struct {
	ID        int64      `gorm:"AUTO_INCREMENT" json:"id"`
	AppUserID int64      `json:appuser_id"`
	Role      string     `json:"role"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at",omitempty"`
}

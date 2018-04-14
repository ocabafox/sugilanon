package models

import "time"

type Story struct {
	ID        int64      `gorm:"AUTO_INCREMENT" json:"id"`
	AppUserID int64      `gorm:"type:varchar(32)" json:app_user_id"`
	Title     string     `json:"title,omitempty"`
	Body      string     `json:"body,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (story *Story) Create() error {
	return db.Debug().Model(&story).Create(&story).Error
}

func (story *Story) Delete() error {
	return db.Debug().Model(&story).Where("id=?", story.ID).Delete(&story).Error
}

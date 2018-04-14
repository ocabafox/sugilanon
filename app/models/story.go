package models

import (
	"encoding/json"
	"time"

	"github.com/XanderDwyl/sugilanon/app/libs/mycache"
)

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

func GetStories() (*[]Story, error) {
	var stories []Story
	key := "stories"
	storiesCache, err := mycache.Get(key)
	if err != nil {
		err = db.Debug().Raw("SELECT * FROM stories").Scan(&stories).Error
		storiesJSON, _ := json.Marshal(&stories)
		_, err := mycache.Set(key, string(storiesJSON), 1800)
		if err != nil {
			return &stories, err
		}
	} else {
		err = json.Unmarshal([]byte(storiesCache), &stories)
	}
	return &stories, err
}

package models

import (
	"encoding/json"
	"time"

	"github.com/XanderDwyl/sugilanon/app/libs/mycache"
)

type Story struct {
	ID        int64      `gorm:"AUTO_INCREMENT" json:"id"`
	AppUserId int64      `json:"appuser_id"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (story *Story) CreateStory() error {
	err := db.Debug().Model(&story).Create(&story).Error

	return err
}

func GetStories() ([]Story, error) {
	var stories []Story
	var err error

	key := "stories"
	cache, err := mycache.Get(key)
	if err != nil {
		err = db.Debug().Model(&Story{}).Scan(&stories).Error

		storiesJSON, _ := json.Marshal(&stories)
		_, err := mycache.Set(key, string(storiesJSON), 1800)
		if err != nil {
			return stories, err
		}
	} else {
		err = json.Unmarshal([]byte(cache), &stories)
	}

	return stories, err
}

package models

import (
	"encoding/json"
	"time"

	"github.com/XanderDwyl/sugilanon/app/libs/mycache"
)

type FacebookAccount struct {
	FacebookId string     `gorm:"type:varchar(32)" json:facebook_id"`
	Name       string     `gorm:"type:varchar(32)" json:"name"`
	Email      string     `gorm:"type:varchar(130)" json:"email"`
	Link       string     `gorm:"type:varchar(130)" json:"link"`
	Gender     string     `gorm:"type:varchar(130)" json:"gender,omitempty"`
	Updated    string     `gorm:"type:varchar(130)" json:"updated"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	DeletedAt  *time.Time `json:"deleted_at",omitempty"`
}

func (fbAccount *FacebookAccount) FacebookCreateUser() (FacebookAccount, error) {
	err := db.Debug().Model(&fbAccount).Create(&fbAccount).Error

	return *fbAccount, err
}

func (fbAccount *FacebookAccount) FacebookUpdateUser() (FacebookAccount, error) {
	err := db.Debug().Model(&fbAccount).Where("facebook_id=?", fbAccount.FacebookId).Omit("facebook_id", "created_at").Updates(&fbAccount).Error

	return *fbAccount, err
}

func (fbAccount *FacebookAccount) GetFacebookAccount() (FacebookAccount, error) {
	var facebookAccount FacebookAccount
	var err error

	key := "facebookAccount"
	cache, err := mycache.Get(key)
	if err != nil {
		err = db.Debug().Model(&FacebookAccount{}).Where("facebook_id=?", fbAccount.FacebookId).Scan(&facebookAccount).Error

		facebookAccountJSON, _ := json.Marshal(facebookAccount)
		_, err := mycache.Set(key, string(facebookAccountJSON), 300)
		if err != nil {
			return facebookAccount, err
		}
	} else {
		err = json.Unmarshal([]byte(cache), facebookAccount)
	}

	return facebookAccount, err
}

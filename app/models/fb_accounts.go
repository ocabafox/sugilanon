package models

import (
	"time"
)

type FacebookAccount struct {
	FacebookId string     `gorm:"type:varchar(32)" json:facebook_id"`
	Name       string     `gorm:"type:varchar(32)" json:"name"`
	Email      string     `gorm:"type:varchar(130)" json:"email"`
	Link       string     `gorm:"type:varchar(130)" json:"link"`
	Gender     string     `gorm:"type:varchar(130)" json:"gender,omitempty"`
	Website    string     `gorm:"type:varchar(130)" json:"website,omitempty"`
	Updated    string     `gorm:"type:varchar(130)" json:"updated"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}

func (fbAccount *FacebookAccount) FacebookCreateUser() (FacebookAccount, error) {
	err := db.Debug().Model(&fbAccount).Create(&fbAccount).Error

	return *fbAccount, err
}

func (fbAccount *FacebookAccount) FacebookUpdateUser() error {
	err := db.Debug().Model(&fbAccount).Omit("facebook_id", "created_at").Updates(&fbAccount).Error

	return err
}

func (fbAccount *FacebookAccount) GetFacebookAccount() (FacebookAccount, error) {
	var facebookAccount FacebookAccount
	err := db.Debug().Model(&FacebookAccount{}).Where("facebook_id=?", fbAccount.FacebookId).Scan(&facebookAccount).Error

	return facebookAccount, err
}

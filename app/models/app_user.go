package models

import (
	"strconv"
	"time"
)

// User ....
type AppUser struct {
	ID                int64      `gorm:"AUTO_INCREMENT" json:"id"`
	FacebookId        string     `gorm:"type:varchar(32)" json:facebook_id"`
	Username          string     `gorm:"type:varchar(32)" json:"username"`
	VerificationToken string     `json:"verification_token,omitempty"`
	IsVerified        bool       `json:"is_verified,omitempty"`
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
	DeletedAt         *time.Time `json:"deleted_at",omitempty"`
}

func (appUser *AppUser) AppUpdate() (AppUser, error) {
	err := db.Debug().Model(&appUser).Omit("facebook_id", "created_at").Updates(&appUser).Error

	return *appUser, err
}

func (appUser *AppUser) AppDelete() error {
	err := db.Debug().Model(&appUser).Where("facebook_id=?", appUser.FacebookId).Delete(&appUser).Error

	return err
}

func AppCreate(facebookId string) (AppUser, error) {
	applicationUser := AppUser{
		FacebookId:        facebookId,
		Username:          "anonymouse" + strconv.Itoa(int(time.Now().UnixNano())),
		VerificationToken: "TOKEN",
		IsVerified:        false,
	}

	err := db.Debug().Model(&AppUser{}).Create(&applicationUser).Error

	return applicationUser, err
}

func GetAppUserById(facebookId string) (AppUser, error) {
	var applicationUser AppUser
	err := db.Debug().Model(&AppUser{}).Where("facebook_id=?", facebookId).Scan(&applicationUser).Error

	return applicationUser, err
}

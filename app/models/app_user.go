package models

import (
	"strconv"
	"time"
)

// User ....
type AppUser struct {
	ID                int64      `gorm:"AUTO_INCREMENT" json:"id"`
	ApplicationId     string     `gorm:"type:varchar(32)" json:application_id"`
	Username          string     `gorm:"type:varchar(32)" json:"username"`
	VerificationToken string     `json:"verification_token,omitempty"`
	IsVerified        bool       `json:"is_verified,omitempty"`
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
	DeletedAt         *time.Time `json:"deleted_at",omitempty"`
}

func (appUser *AppUser) AppUpdate() (AppUser, error) {
	err := db.Debug().Model(&appUser).Omit("application_id", "created_at").Updates(&appUser).Error

	return *appUser, err
}

func (appUser *AppUser) AppDelete() error {
	err := db.Debug().Model(&appUser).Where("application_id=?", appUser.ApplicationId).Delete(&appUser).Error

	return err
}

func AppCreate(applicationId string) (AppUser, error) {
	applicationUser := AppUser{
		ApplicationId:     applicationId,
		Username:          "anonymouse" + strconv.Itoa(int(time.Now().UnixNano())),
		VerificationToken: "TOKEN",
		IsVerified:        false,
	}

	err := db.Debug().Model(&AppUser{}).Create(&applicationUser).Error

	return applicationUser, err
}

func GetAppUserById(applicationId string) (AppUser, error) {
	var applicationUser AppUser
	err := db.Debug().Model(&AppUser{}).Where("application_id=?", applicationId).Scan(&applicationUser).Error

	return applicationUser, err
}

func GetAppUserByUsername(appUsername string) (AppUser, error) {
	var applicationUser AppUser
	err := db.Debug().Model(&AppUser{}).Where("username=?", appUsername).Scan(&applicationUser).Error

	return applicationUser, err
}

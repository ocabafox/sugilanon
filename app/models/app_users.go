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

type AppUserByRole struct {
	IsVerified bool   `json:"is_verified"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Link       string `json:"link"`
	Role       string `json:"role"`
}

func (appUser *AppUser) AppUpdateUser() (AppUser, error) {
	err := db.Debug().Model(&appUser).Omit("application_id", "created_at").Updates(&appUser).Error

	return *appUser, err
}

func (appUser *AppUser) AppDeleteUser() error {
	err := db.Debug().Model(&appUser).Where("application_id=?", appUser.ApplicationId).Delete(&appUser).Error

	return err
}

func AppCreateUser(applicationId string) (AppUser, error) {
	applicationUser := AppUser{
		ApplicationId:     applicationId,
		Username:          "anonymouse" + strconv.Itoa(int(time.Now().UnixNano())),
		VerificationToken: "TOKEN",
		IsVerified:        false,
	}

	err := db.Debug().Model(&AppUser{}).Create(&applicationUser).Error

	return applicationUser, err
}

func GetAppUserByFacebookId(applicationId string) (AppUser, error) {
	var applicationUser AppUser
	err := db.Debug().Model(&AppUser{}).Where("application_id=?", applicationId).Scan(&applicationUser).Error

	return applicationUser, err
}

func GetAppUserByUsername(appUsername string) (AppUser, error) {
	var applicationUser AppUser
	err := db.Debug().Model(&AppUser{}).Where("username=?", appUsername).Scan(&applicationUser).Error

	return applicationUser, err
}

func GetAppUsers() ([]AppUser, error) {
	var applicationUsers []AppUser
	err := db.Debug().Model(&AppUser{}).Scan(&applicationUsers).Error

	return applicationUsers, err
}

func GetAppUsersByRole() ([]AppUserByRole, error) {
	var applicationUsersByRole []AppUserByRole
	err := db.Debug().Table("facebook_accounts").Select("facebook_accounts.name, facebook_accounts.email, facebook_accounts.link, app_users.username, app_users.is_verified, app_user_roles.role").Joins("INNER JOIN app_users ON facebook_accounts.facebook_id = app_users.application_id").Joins("INNER JOIN app_user_roles ON app_users.id = app_user_roles.app_user_id").Scan(&applicationUsersByRole).Error

	return applicationUsersByRole, err
}

package models

import (
	"encoding/json"
	"time"

	"github.com/XanderDwyl/sugilanon/app/libs/mycache"
)

type AppUserRole struct {
	ID        int64      `gorm:"AUTO_INCREMENT" json:"id"`
	AppUserId int64      `json:"appuser_id"`
	Role      string     `json:"role"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at",omitempty"`
}

func AppCreateUserRole(appUserId int64, role string) error {
	applicationUserRole := AppUserRole{
		AppUserId: appUserId,
		Role:      role,
	}

	err := db.Debug().Model(&AppUserRole{}).Create(&applicationUserRole).Error

	return err
}

func AppDeleteUserRole(appUserId int64) error {
	var applicationUserRole AppUserRole
	err := db.Debug().Model(&AppUserRole{}).Where("app_user_id=?", appUserId).Delete(&applicationUserRole).Error

	return err
}

func GetAppUserRoleByAppUserId(appUserId int64) (AppUserRole, error) {
	var applicationUserRole AppUserRole
	var err error

	key := "applicationUserRole"
	cache, err := mycache.Get(key)
	if err != nil {
		err = db.Debug().Model(&AppUserRole{}).Where("app_user_id=?", appUserId).Scan(&applicationUserRole).Error

		applicationUserRoleJSON, _ := json.Marshal(applicationUserRole)
		_, err := mycache.Set(key, string(applicationUserRoleJSON), 1800)
		if err != nil {
			return applicationUserRole, err
		}
	} else {
		err = json.Unmarshal([]byte(cache), &applicationUserRole)
	}

	return applicationUserRole, err
}

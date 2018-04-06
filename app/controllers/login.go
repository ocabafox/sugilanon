package controllers

import (
	"github.com/XanderDwyl/sugilanon/app/models"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	account := models.FacebookAccount{
		FacebookId: c.PostForm("facebook_id"),
		Name:       c.PostForm("name"),
		Email:      c.PostForm("email"),
		Link:       c.PostForm("link"),
		Gender:     c.PostForm("gender"),
		Website:    c.PostForm("website"),
		Updated:    c.PostForm("updated"),
	}

	appUser := models.AppUser{}
	fbUser, err := account.GetFacebookAccount()
	if err != nil {
		fbUser, _ = account.FacebookCreateUser()
		appUser, _ = models.AppCreateUser(account.FacebookId)
		appUsers, _ := models.GetAppUsers()
		if len(appUsers) == 1 {
			models.AppCreateUserRole(appUser.ID, "admin")
		} else {
			models.AppCreateUserRole(appUser.ID, "user")
		}
	} else {
		appUser, err = models.GetAppUserById(fbUser.FacebookId)
		if err != nil {
			appUser, _ = models.AppCreateUser(fbUser.FacebookId)

			models.AppCreateUserRole(appUser.ID, "user")
		}

		if account.Updated != fbUser.Updated {
			fbUser, _ = account.FacebookUpdateUser()
		}
	}

	user := User{
		IsVerified: appUser.IsVerified,
		Username:   appUser.Username,
		FacebookId: fbUser.FacebookId,
		Name:       fbUser.Name,
		Email:      fbUser.Email,
	}

	SetAuth(c, user)
}

func Logout(c *gin.Context) {
	ClearAuth(c)
}

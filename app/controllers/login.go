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

	applicationUser := models.AppUser{}
	facebookAccount, err := account.GetFacebookAccount()
	if err != nil {
		facebookAccount, _ = account.FacebookCreate()
		applicationUser, _ = models.AppCreate(account.FacebookId)
	} else {
		applicationUser, err = models.GetAppUserById(facebookAccount.FacebookId)
		if err != nil {
			applicationUser, _ = models.AppCreate(facebookAccount.FacebookId)
		}

		if account.Updated == facebookAccount.Updated {
			account = facebookAccount
		} else {
			account.FacebookUpdate()
		}
	}

	user := User{
		IsVerified: applicationUser.IsVerified,
		Username:   applicationUser.Username,
		FacebookId: facebookAccount.FacebookId,
		Name:       facebookAccount.Name,
		Email:      facebookAccount.Email,
	}

	SetAuth(c, user)
}

func Logout(c *gin.Context) {
	ClearAuth(c)
}

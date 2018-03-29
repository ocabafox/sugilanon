package controllers

import (
	"github.com/XanderDwyl/sugilanon/app/models"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	account := models.User{
		FacebookId: c.PostForm("facebook_id"),
		Name:       c.PostForm("name"),
		Email:      c.PostForm("email"),
		Link:       c.PostForm("link"),
		Gender:     c.PostForm("gender"),
		Website:    c.PostForm("website"),
		Updated:    c.PostForm("updated"),
	}

	user, _ := account.GetUser()
	if user.ID == 0 {
		account, _ = account.Create()
	} else {
		if account.Updated == user.Updated {
			account = user
		} else {
			account, _ = account.Update()
		}
	}

	SetAuth(c, account)
}

func Logout(c *gin.Context) {
	ClearAuth(c)
}

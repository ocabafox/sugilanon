package controllers

import (
	"github.com/XanderDwyl/sugilanon/app/models"
	"github.com/gin-gonic/gin"
)

func LoginIndex(c *gin.Context) {
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
	if len(user) == 0 {
		account, _ = account.Create()
	} else {
		if account.Updated == user[0].Updated {
			account = user[0]
		} else {
			account, _ = account.Update()
		}
	}

	SetAuth(c, account)
}

func LogoutIndex(c *gin.Context) {
	ClearAuth(c)
}

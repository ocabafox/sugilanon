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
	}

	user, _ := account.GetUser()
	if len(user) == 0 {
		account, _ = account.Create()
	} else {
		account = models.User{
			FacebookId: user[0].FacebookId,
			Username:   user[0].Username,
			Name:       user[0].Name,
			Email:      user[0].Email,
			IsVerified: user[0].IsVerified,
		}
	}

	SetAuth(c, account)
}

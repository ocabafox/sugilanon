package controllers

import (
	"github.com/XanderDwyl/sugilanon/app/models"
	"github.com/gin-gonic/gin"
)

func ProfileIndex(c *gin.Context) {
	if !IsLogin(c) {
		c.Redirect(302, "/")
		c.Abort()

		return
	}

	appUser, err := models.GetAppUserById(GetFacebookId(c))
	if err != nil {
		c.Redirect(302, "/")
		c.Abort()

		return
	}

	RenderHTML(c, gin.H{
		"page":    "PROFILE",
		"appUser": appUser,
	})
}

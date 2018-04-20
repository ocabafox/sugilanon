package controllers

import (
	"log"
	"github.com/XanderDwyl/sugilanon/app/models"
	"github.com/gin-gonic/gin"
)

func ProfileIndex(c *gin.Context) {
	if !IsLogin(c) {
		log.Println("=================")
		log.Println("not login")
		c.Redirect(302, "/")
		c.Abort()

		return
	}

	appUser, err := models.GetAppUserByFacebookId(GetFacebookId(c))
	if err != nil {
		log.Println("=================")
		log.Println("fb user not found")
		log.Println(GetFacebookId(c))
		log.Println("=================")
		c.Redirect(302, "/")
		c.Abort()

		return
	}

	RenderHTML(c, gin.H{
		"page":    "PROFILE",
		"appUser": appUser,
	})
}

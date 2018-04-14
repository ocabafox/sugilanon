package controllers

import (
	"log"

	"github.com/XanderDwyl/sugilanon/app/models"
	"github.com/gin-gonic/gin"
)

func AdminUsersIndex(c *gin.Context) {
	fbAccounts, err := models.GetFacebookAccounts()
	if err != nil {
		c.Redirect(302, "/admin")
		c.Abort()

		return
	}

	log.Println(fbAccounts)

	RenderHTML(c, gin.H{
		"page":       "USERS | ADMIN",
		"fbAccounts": fbAccounts,
	})
}

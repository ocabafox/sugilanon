package controllers

import (
	"github.com/XanderDwyl/sugilanon/app/models"
	"github.com/gin-gonic/gin"
)

func AdminIndex(c *gin.Context) {
	appUserRole, err := models.GetAppUserRoleByAppUserId(GetAppUserId(c))
	if err != nil || appUserRole.Role != "admin" {
		c.Redirect(302, "/")
		c.Abort()

		return
	}

	RenderHTML(c, gin.H{
		"page": "ADMIN",
	})
}

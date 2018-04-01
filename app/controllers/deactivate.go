package controllers

import (
	"github.com/XanderDwyl/sugilanon/app/models"
	"github.com/gin-gonic/gin"
)

// Deactivate ...
func Deactivate(c *gin.Context) {
	if !IsLogin(c) {
		c.Redirect(302, "/")
		c.Abort()

		return
	}

	appUsername := c.Param("username")
	appUser, err := models.GetAppUserByUsername(appUsername)
	if err != nil {
		c.Redirect(302, "/profile")
		c.Abort()

		return
	}

	ClearAuth(c)
	appUser.AppDelete()

	c.Redirect(302, "/")
	c.Abort()
}

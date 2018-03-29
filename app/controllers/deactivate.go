package controllers

import (
	"github.com/XanderDwyl/sugilanon/app/models"
	"github.com/gin-gonic/gin"
)

func Deactivate(c *gin.Context) {
	if !IsLogin(c) {
		c.Redirect(302, "/")
		c.Abort()

		return
	}

	facebookId := c.Param("facebook_id")
	user, err := models.GetUserById(facebookId)
	if err != nil {
		c.Redirect(302, "/")
		c.Abort()

		return
	}

	user.Delete()
	ClearAuth(c)

	c.Redirect(302, "/")
	c.Abort()
}

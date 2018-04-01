package controllers

import (
	"log"

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
	log.Println(appUsername)
	//appUser, err := models.GetAppUserById(facebookID)
	//if err != nil {
	//c.Redirect(302, "/")
	//c.Abort()

	//return
	//}

	//ClearAuth(c)
	//appUser.AppDelete()

	c.Redirect(302, "/")
	c.Abort()
}

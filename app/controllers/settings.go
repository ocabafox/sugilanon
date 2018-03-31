package controllers

import (
	"github.com/gin-gonic/gin"
)

func SettingsIndex(c *gin.Context) {
	if !IsLogin(c) {
		c.Redirect(302, "/")
		c.Abort()

		return
	}

	RenderHTML(c, gin.H{})
}

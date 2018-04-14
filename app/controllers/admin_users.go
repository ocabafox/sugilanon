package controllers

import (
	"github.com/gin-gonic/gin"
)

func AdminUsersIndex(c *gin.Context) {
	RenderHTML(c, gin.H{
		"page": "USERS | ADMIN",
	})
}

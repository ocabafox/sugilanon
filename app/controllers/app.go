package controllers

import (
	"github.com/gin-gonic/gin"
)

// AppIndex ...
func AppIndex(c *gin.Context) {
	RenderHTML(c, gin.H{
		"page": "HOME",
	})
}

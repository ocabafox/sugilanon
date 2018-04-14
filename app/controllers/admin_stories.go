package controllers

import (
	"github.com/gin-gonic/gin"
)

func AdminStoriesIndex(c *gin.Context) {
	RenderHTML(c, gin.H{
		"page": "STORIES | ADMIN",
	})
}

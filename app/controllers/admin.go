package controllers

import (
	"github.com/gin-gonic/gin"
)

func AdminIndex(c *gin.Context) {
	RenderHTML(c, gin.H{
		"page": "QUESTIONS | ADMIN",
	})
}

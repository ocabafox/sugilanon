package controllers

import (
	"github.com/gin-gonic/gin"
)

func AboutIndex(c *gin.Context) {
	RenderHTML(c, gin.H{
		"page": "ABOUT US",
	})
}

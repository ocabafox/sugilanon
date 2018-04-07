package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
)

// AppIndex ...
func AppIndex(c *gin.Context) {
	var fbApplicationId string

	if os.Getenv("MODE") != "production" {
		fbApplicationId = "1692775630777637"
	} else {
		fbApplicationId = "1618996761552073"
	}

	RenderHTML(c, gin.H{
		"page":            "HOME",
		"fbApplicationId": fbApplicationId,
	})
}

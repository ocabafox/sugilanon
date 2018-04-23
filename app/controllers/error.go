package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
)

// PageNotFound ...
func PageNotFound(c *gin.Context) {
	log.Print("PageNotFound")

	RenderHTML(c, gin.H{
		"page": "PAGE NOT FOUND",
	})
}

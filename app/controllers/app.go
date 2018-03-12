package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
)

// AppIndex ...
func AppIndex(c *gin.Context) {
	log.Println("Welcome to Sugilanon!")
	RenderHTML(c, gin.H{})
}

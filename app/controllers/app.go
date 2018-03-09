package controllers

import (
	"log"

	"gopkg.in/gin-gonic/gin.v1"
)

// AppIndex ...
func AppIndex(c *gin.Context) {
	log.Println("Welcome to Sugilanon!")
	RenderHTML(c, gin.H{})
}

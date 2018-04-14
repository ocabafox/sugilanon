package controllers

import (
	"log"

	"github.com/XanderDwyl/sugilanon/app/models"
	"github.com/gin-gonic/gin"
)

func Story(c *gin.Context) {
	title := c.PostForm("title")
	body := c.PostForm("body")

	story := models.Story{}
	story.AppUserID = GetAppUserId(c)
	story.Title = title
	story.Body = body

	err := story.Create()

	if err != nil {
		log.Println(err)
		return
	}
	NoRoute(c)
}

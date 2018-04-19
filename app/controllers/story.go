package controllers

import (
	"strconv"

	"github.com/XanderDwyl/sugilanon/app/models"
	"github.com/gin-gonic/gin"
)

func Story(c *gin.Context) {
	owner := c.PostForm("owner")
	appUserId, _ := strconv.ParseInt(owner, 10, 64)
	story := models.Story{
		AppUserId: appUserId,
		Title:     c.PostForm("title"),
		Body:      c.PostForm("body"),
	}

	err := story.CreateStory()
	if err != nil {
		OutputJSON(c, "error", "Unable to create story")

		return
	}

	OutputJSON(c, "success", "Story successfully created")
}

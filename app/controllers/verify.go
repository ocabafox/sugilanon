package controllers

import (
	"log"

	"github.com/XanderDwyl/sugilanon/app/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func VerifyIndex(c *gin.Context) {
	log.Println("VERIFY")

	facebookId := c.Param("facebook_id")
	verificationToken := c.Param("verification_token")

	successfulVerification := true
	user, err := models.GetUserById(facebookId)
	if err != nil || user.VerificationToken != verificationToken {
		successfulVerification = false
	}

	user.IsVerified = true
	session := sessions.Default(c)
	session.Set("is_verified", user.IsVerified)
	user.Update()

	RenderHTML(c, gin.H{
		"successfulVerification": successfulVerification,
	})
}

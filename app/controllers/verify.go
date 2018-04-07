package controllers

import (
	"github.com/XanderDwyl/sugilanon/app/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func VerifyIndex(c *gin.Context) {
	facebookId := c.Param("facebook_id")
	verificationToken := c.Param("verification_token")

	successfulVerification := false
	appUser, _ := models.GetAppUserByFacebookId(facebookId)
	if !appUser.IsVerified && appUser.VerificationToken == verificationToken {
		successfulVerification = true

		session := sessions.Default(c)
		session.Set("is_verified", successfulVerification)
		session.Save()

		appUser.IsVerified = successfulVerification
		appUser.AppUpdateUser()
	}

	RenderHTML(c, gin.H{
		"successfulVerification": successfulVerification,
		"page": "VERIFICATION PAGE",
	})
}

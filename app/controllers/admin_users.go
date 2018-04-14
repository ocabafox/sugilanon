package controllers

import (
	"github.com/XanderDwyl/sugilanon/app/models"
	"github.com/gin-gonic/gin"
)

func AdminUsersIndex(c *gin.Context) {
	appUsers, err := models.GetAppUsers()
	if err != nil {
		c.Redirect(302, "/admin")
		c.Abort()

		return
	}

	var admins []User
	var users []User
	for _, appUsersValue := range appUsers {
		fbUser, err := models.GetFacebookAccountByFacebookId(appUsersValue.ApplicationId)
		if err != nil {
			c.Redirect(302, "/admin")
			c.Abort()

			return
		}

		appUserRole, err := models.GetAppUserRoleByAppUserId(appUsersValue.ID)
		if err != nil {
			c.Redirect(302, "/admin")
			c.Abort()

			return
		}

		user := User{
			IsVerified: appUsersValue.IsVerified,
			Username:   appUsersValue.Username,
			Name:       fbUser.Name,
			Email:      fbUser.Email,
			Link:       fbUser.Link,
		}

		if appUserRole.Role == "admin" {
			admins = append(admins, user)
		} else {
			users = append(users, user)
		}
	}

	RenderHTML(c, gin.H{
		"page":   "USERS | ADMIN",
		"admins": admins,
		"users":  users,
	})
}

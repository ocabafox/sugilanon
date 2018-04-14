package controllers

import (
	"github.com/XanderDwyl/sugilanon/app/models"
	"github.com/gin-gonic/gin"
)

func AdminUsersIndex(c *gin.Context) {
	appUsersByRole, err := models.GetAppUsersByRole()
	if err != nil {
		c.Redirect(302, "/admin")
		c.Abort()

		return
	}

	var admins []User
	var users []User
	for _, usersValue := range appUsersByRole {
		if usersValue.Role == "admin" {
			admins = append(admins, User{
				IsVerified: usersValue.IsVerified,
				Username:   usersValue.Username,
				Name:       usersValue.Name,
				Email:      usersValue.Email,
				Link:       usersValue.Link,
			})
		} else {
			users = append(users, User{
				IsVerified: usersValue.IsVerified,
				Username:   usersValue.Username,
				Name:       usersValue.Name,
				Email:      usersValue.Email,
				Link:       usersValue.Link,
			})
		}
	}

	RenderHTML(c, gin.H{
		"page":   "USERS | ADMIN",
		"admins": admins,
		"users":  users,
	})
}

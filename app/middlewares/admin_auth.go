package middleware

import (
	"github.com/XanderDwyl/sugilanon/app/controllers"
	"github.com/XanderDwyl/sugilanon/app/models"
	"github.com/gin-gonic/gin"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		appUserRole, err := models.GetAppUserRoleByAppUserId(controllers.GetAppUserId(c))
		if err != nil || appUserRole.Role != "admin" {
			c.Redirect(302, "/")
			c.Abort()
		}
	}
}

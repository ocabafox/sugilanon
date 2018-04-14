package main

import (
	"os"
	"time"

	"github.com/XanderDwyl/sugilanon/app/controllers"
	"github.com/XanderDwyl/sugilanon/app/libs/ezgintemplate"
	"github.com/XanderDwyl/sugilanon/app/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	models.Setup()
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	router := gin.Default()
	allowOrigins := []string{
		"https://sugilanon.com",
		"https://www.sugilanon.com",
	}

	if os.Getenv("MODE") != "production" {
		allowOrigins = append(allowOrigins, "https://localhost:3000")
	}

	store := sessions.NewCookieStore([]byte("Lod5c5F"))
	router.Static("/assets", "./assets")
	router.Use(sessions.Sessions("mysession", store))
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Access-Control-Allow-Origin", "Accept", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	render := ezgintemplate.New()
	render.TemplatesDir = "app/views/"
	render.Layout = "layouts/base"
	render.Ext = ".tmpl"
	render.Debug = true
	router.HTMLRender = render.Init()

	initializeRoutes(router)

	if os.Getenv("MODE") != "production" {
		router.RunTLS(":"+port, "./certificate/server.crt", "./certificate/server.key")
	} else {
		router.Run(":" + port)
	}
}

func initializeRoutes(origRouter *gin.Engine) {
	origRouter.NoRoute(controllers.NoRoute)

	router := origRouter.Group("")
	{
		router.GET("/", controllers.AppIndex)
		router.GET("/about", controllers.AboutIndex)
		router.GET("/profile", controllers.ProfileIndex)
		router.GET("/logout", controllers.Logout)
		router.GET("/verify/:facebook_id/:verification_token", controllers.VerifyIndex)

		router.POST("/login", controllers.Login)
		router.POST("/deactivate/:username", controllers.Deactivate)
	}

	admin := origRouter.Group("/admin")
	{
		admin.GET("/", controllers.AdminIndex)
		admin.GET("/stories", controllers.AdminStoriesIndex)
		admin.GET("/users", controllers.AdminUsersIndex)
	}
}

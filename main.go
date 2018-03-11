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
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"gopkg.in/gin-gonic/gin.v1"
)

func init() {
	models.Setup()
}

func main() {
	router := gin.Default()
	allowOrigins := []string{
		"http://sugilanon.com",
		"https://sugilanon.com",
	}

	if os.Getenv("MODE") != "production" {
		allowOrigins = append(allowOrigins, "http://localhost:8080")
		allowOrigins = append(allowOrigins, "http://localhost:8081")
		allowOrigins = append(allowOrigins, "http://127.0.0.1:8080")
		allowOrigins = append(allowOrigins, "http://127.0.0.1:8081")
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Access-Control-Allow-Origin", "Accept", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	store := sessions.NewCookieStore([]byte("Lod5c5F"))

	router.Use(sessions.Sessions("mysession", store))
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Static("/assets", "./assets")
	router.Use(gin.Recovery())

	render := ezgintemplate.New()
	render.TemplatesDir = "app/views/"
	render.Layout = "layouts/base"
	render.Ext = ".tmpl"
	render.Debug = true

	router.HTMLRender = render.Init()
	initializeRoutes(router)

	router.Run(":3000")
}
func initializeRoutes(origRouter *gin.Engine) {
	router := origRouter.Group("")

	router.GET("/", controllers.AppIndex)
}

package controllers

import (
	"os"
	"runtime"
	"strings"

	"github.com/XanderDwyl/sugilanon/app/libs/tmplname"
	"gopkg.in/gin-gonic/gin.v1"
)

// OutputJSON ...
func OutputJSON(c *gin.Context, status, msg string) {
	c.JSON(200, gin.H{
		"status":  status,
		"message": msg,
	})
}

// OutputDataJSON ...
func OutputDataJSON(c *gin.Context, status, msg string, data gin.H) {
	c.JSON(200, gin.H{
		"status":  status,
		"message": msg,
		"data":    data,
	})
}

// RenderHTML ...
func RenderHTML(c *gin.Context, data gin.H) {

	pc, _, _, _ := runtime.Caller(1)
	callerName := runtime.FuncForPC(pc).Name()
	for strings.Contains(callerName, ".") {
		a := strings.Index(callerName, ".")
		callerName = callerName[a+1:]
	}

	tmpl := tmplname.Convert(callerName)
	// check whether the file is existed.
	_, err := os.Stat("app/views/" + tmpl + ".tmpl")
	if err != nil {
		c.String(200, "%s not found", "app/views/"+tmpl+".tmpl")
		return
	}

	RenderTemplate(c, tmpl, data, 200)
}

// RenderTemplate ...
func RenderTemplate(c *gin.Context, tmpl string, data gin.H, statusCode int) {
	data["host"] = c.Request.Host

	c.HTML(statusCode, tmpl, data)
}

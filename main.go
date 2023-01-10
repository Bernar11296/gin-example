package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", showIndexPage)

	router.GET("/article/view/:article_id", getArticle)

	router.Run()
}

func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(200, data["payload"])
	case "application/xml":
		c.XML(200, data["payload"])
	default:
		c.HTML(200, templateName, data)
	}
}

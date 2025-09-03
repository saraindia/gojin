package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Welcome",
	})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", HomePageHandler)
	router.Run(":80")
}

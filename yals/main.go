package main

import (
	"yals/server"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.StaticFile("/", "./server/static/_index.html")
	r.Static("/static", "./server/static")

	// r.GET("/", server.HandleIndex)
	r.POST("/add", server.HandlePostRequest)
	r.GET("/go/:Shortcut", server.HandleGetRequest)
	r.Run()
}

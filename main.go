package main

import (
	"restapi/controllers"
	"restapi/models" // new

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks) // new
	r.Run()
}

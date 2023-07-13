package main

import (
	"go-gin/controllers"
	"go-gin/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/post", controllers.PostsCreate)
	r.GET("/post/:id", controllers.PostsShow)
	r.PUT("/post/:id", controllers.PostsUpdate)
	r.DELETE("/post/:id", controllers.PostsDelete)
	r.GET("/post", controllers.PostsIndex)

	r.Run() // listen and serve on 0.0.0.0:8080
}

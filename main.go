package main

import (
	repository "gomongo/Repository"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/posts", repository.GetPost)
	r.POST("/posts", repository.SendPost)
	r.GET("/posts/:id", repository.ShowPost)
	r.PUT("/posts/:id", repository.UpdatePost)
	r.DELETE("/posts/:id", repository.DeletePost)

	r.Run(":8080")
}

package main

import (
	repository "gomongo/Repository"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/posts", repository.GetPost)

	r.Run(":8080")
}

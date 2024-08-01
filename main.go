package main

import (
	"docker/app/controller/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/add/data", controller.Add)
	r.GET("/data/list", controller.Get)

	r.Run(":5544")
}
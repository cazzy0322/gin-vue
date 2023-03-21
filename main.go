package main

import (
	"github.com/cazzy0322/gin-vue/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	r.Run(":8080")
}

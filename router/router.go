package router

import (
	"github.com/cazzy0322/gin-vue/controller"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	r.GET("/api/user/register", controller.Register)

	r.Run(":8080")
}

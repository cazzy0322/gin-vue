package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shuwenhe/utils"
	"net/http"
)

type Student struct {
	Name  string
	Age   int
	Skill string
}

func Register(c *gin.Context) {
	// Get the parameter
	//name := c.PostForm("name")
	password := c.PostForm("password")
	phone := c.PostForm("phone")

	// Data verification
	if len(phone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "The phone num must be 11 digits!",
		})
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "password cannot be less than 6 digits!",
		})
	}
	// Create user

	// Return result

	c.JSON(utils.NewSucc("Register success!", gin.H{
		"msg": "Register success",
	}))
}

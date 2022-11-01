package api

import (
	"github.com/gin-gonic/gin"
	"user/handler"
	"user/pkg/e"
)

func UserLogin(c *gin.Context) {

}

func UserRegister(c *gin.Context) {
	var userRegister handler.UserService

	if err := c.ShouldBind(&userRegister); err != nil {
		c.JSON(e.InvalidParams, err)
	}

	res := userRegister.RegisterHandler()
	c.JSON(e.SUCCESS, res)
}

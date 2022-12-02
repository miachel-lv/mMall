package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user/handler"
)

func UserLogin(c *gin.Context) {
	var userRegister handler.UserService

	if err := c.ShouldBind(&userRegister); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	res := userRegister.LoginHandler(c)
	c.JSON(http.StatusOK, res)
}

func UserRegister(c *gin.Context) {
	var userRegister handler.UserService

	if err := c.ShouldBind(&userRegister); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	res := userRegister.RegisterHandler()
	c.JSON(http.StatusOK, res)
}

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "index",
	})
}
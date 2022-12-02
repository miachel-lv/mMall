package api

import (
	"github.com/gin-gonic/gin"
	"user/dao"
	"user/handler"
	"user/pkg/e"
	"user/pkg/utils"
)

func CreateProduct(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	createProductService := handler.ProductService{}
	//c.SaveUploadedFile()
	if err := c.ShouldBind(&createProductService); err == nil {
		
	}
}

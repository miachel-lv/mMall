package api

import (
	"github.com/gin-gonic/gin"
	"user/handler"
	"user/pkg/utils"
)

func CreateProduct(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	createProductService := handler.ProductService{}
	//c.SaveUploadedFile()
	if err := c.ShouldBind(&createProductService); err == nil {
		res := createProductService.Create(claim.ID, files)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

func UpdateProduct(c *gin.Context) {
	updateProductService := handler.ProductService{}
	if err := c.ShouldBind(&updateProductService); err == nil {
		res := updateProductService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		//utils.LogrusObj.Infoln(err)
	}
}

func DeleteProduct(c *gin.Context) {
	deleteProductService := handler.ProductService{}
	res := deleteProductService.Delete(c.Param("id"))
	c.JSON(200, res)
}


func ListProducts(c *gin.Context) {
	listProductsService := handler.ProductService{}
	if err := c.ShouldBind(&listProductsService); err == nil {
		res := listProductsService.List()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		//util.LogrusObj.Infoln(err)
	}
}

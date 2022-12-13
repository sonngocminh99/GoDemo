package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sonngocminh99/GoDemo/dto"
	"github.com/sonngocminh99/GoDemo/entity"
	"github.com/sonngocminh99/GoDemo/helper"
	"github.com/sonngocminh99/GoDemo/service"
)

type ProductController interface {
	AllProducts(context *gin.Context)
	FindProductById(context *gin.Context)
	UpdateProduct(context *gin.Context)
	DeleteProduct(context *gin.Context)
	InsertProduct(context *gin.Context)
}

type productController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &productController{
		productService: productService}
}

func (c *productController) AllProducts(context *gin.Context) {
	var products []entity.Product = c.productService.AllProducts()
	res := helper.BuildResponse(true, "OK", products)
	context.JSON(http.StatusOK, res)
}

func (c *productController) FindProductById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildError("No param id was found in context", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var product entity.Product = c.productService.FindProductById(int(id))
	if (product == entity.Product{}) {
		res := helper.BuildError("Data not found", "No data with id given", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", product)
		context.JSON(http.StatusOK, res)
	}
}

func (c *productController) UpdateProduct(context *gin.Context) {

	var productUpdateDTO dto.ProductUpdateDTO
	errDTO := context.ShouldBind(&productUpdateDTO)
	if errDTO != nil {
		res := helper.BuildError("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	res := c.productService.UpdateProduct(productUpdateDTO)
	response := helper.BuildResponse(true, "OK", res)
	context.JSON(http.StatusOK, response)

}

func (c *productController) DeleteProduct(context *gin.Context) {
	var product entity.Product
	id, err := strconv.ParseInt(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildError("No param id was found in context", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	product.ID = int(id)
	c.productService.DeleteProduct(product)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)

}

func (c *productController) InsertProduct(context *gin.Context) {

	var productCreateDTO dto.ProductCreateDTO
	errDTO := context.ShouldBind(&productCreateDTO)
	if errDTO != nil {
		res := helper.BuildError("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		result := c.productService.InsertProduct(productCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	db *gorm.DB
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{db: db}
}

func (ctrl *ProductController) GetStoreProducts(c *gin.Context) {
	// TODO: Implement get store products (public)
	c.JSON(http.StatusOK, gin.H{"message": "Get store products - TODO"})
}

func (ctrl *ProductController) GetStoreProduct(c *gin.Context) {
	// TODO: Implement get store product by slug (public)
	c.JSON(http.StatusOK, gin.H{"message": "Get store product - TODO"})
}

func (ctrl *ProductController) GetProducts(c *gin.Context) {
	// TODO: Implement get products for store owner
	c.JSON(http.StatusOK, gin.H{"message": "Get products - TODO"})
}

func (ctrl *ProductController) CreateProduct(c *gin.Context) {
	// TODO: Implement create product
	c.JSON(http.StatusOK, gin.H{"message": "Create product - TODO"})
}

func (ctrl *ProductController) UpdateProduct(c *gin.Context) {
	// TODO: Implement update product
	c.JSON(http.StatusOK, gin.H{"message": "Update product - TODO"})
}

func (ctrl *ProductController) DeleteProduct(c *gin.Context) {
	// TODO: Implement delete product
	c.JSON(http.StatusOK, gin.H{"message": "Delete product - TODO"})
}

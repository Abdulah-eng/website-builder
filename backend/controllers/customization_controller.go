package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CustomizationController struct {
	db *gorm.DB
}

func NewCustomizationController(db *gorm.DB) *CustomizationController {
	return &CustomizationController{db: db}
}

func (ctrl *CustomizationController) GetStorePages(c *gin.Context) {
	// TODO: Implement get store pages
	c.JSON(http.StatusOK, gin.H{"message": "Get store pages - TODO"})
}

func (ctrl *CustomizationController) GetStorePage(c *gin.Context) {
	// TODO: Implement get store page by slug
	c.JSON(http.StatusOK, gin.H{"message": "Get store page - TODO"})
}

func (ctrl *CustomizationController) CreatePage(c *gin.Context) {
	// TODO: Implement create page
	c.JSON(http.StatusOK, gin.H{"message": "Create page - TODO"})
}

func (ctrl *CustomizationController) UpdatePage(c *gin.Context) {
	// TODO: Implement update page
	c.JSON(http.StatusOK, gin.H{"message": "Update page - TODO"})
}

func (ctrl *CustomizationController) DeletePage(c *gin.Context) {
	// TODO: Implement delete page
	c.JSON(http.StatusOK, gin.H{"message": "Delete page - TODO"})
}

func (ctrl *CustomizationController) GetStoreSettings(c *gin.Context) {
	// TODO: Implement get store settings
	c.JSON(http.StatusOK, gin.H{"message": "Get store settings - TODO"})
}

func (ctrl *CustomizationController) UpdateStoreSettings(c *gin.Context) {
	// TODO: Implement update store settings
	c.JSON(http.StatusOK, gin.H{"message": "Update store settings - TODO"})
}

func (ctrl *CustomizationController) GetStoreTheme(c *gin.Context) {
	// TODO: Implement get store theme
	c.JSON(http.StatusOK, gin.H{"message": "Get store theme - TODO"})
}

func (ctrl *CustomizationController) UpdateStoreTheme(c *gin.Context) {
	// TODO: Implement update store theme
	c.JSON(http.StatusOK, gin.H{"message": "Update store theme - TODO"})
}

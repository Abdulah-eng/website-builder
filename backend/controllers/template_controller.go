package controllers

import (
	"net/http"
	"strconv"

	"storemaker-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TemplateController struct {
	db *gorm.DB
}

func NewTemplateController(db *gorm.DB) *TemplateController {
	return &TemplateController{db: db}
}

func (ctrl *TemplateController) GetPublicTemplates(c *gin.Context) {
	var templates []models.Template
	if err := ctrl.db.Where("is_active = ?", true).Find(&templates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch templates"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": templates})
}

func (ctrl *TemplateController) GetTemplate(c *gin.Context) {
	templateID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid template ID"})
		return
	}

	var template models.Template
	if err := ctrl.db.Where("id = ? AND is_active = ?", templateID, true).First(&template).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Template not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch template"})
		}
		return
	}

	c.JSON(http.StatusOK, template)
}

func (ctrl *TemplateController) GetAllTemplates(c *gin.Context) {
	// TODO: Implement get all templates (admin only)
	c.JSON(http.StatusOK, gin.H{"message": "Get all templates - TODO"})
}

func (ctrl *TemplateController) CreateTemplate(c *gin.Context) {
	// TODO: Implement create template (admin only)
	c.JSON(http.StatusOK, gin.H{"message": "Create template - TODO"})
}

func (ctrl *TemplateController) UpdateTemplate(c *gin.Context) {
	// TODO: Implement update template (admin only)
	c.JSON(http.StatusOK, gin.H{"message": "Update template - TODO"})
}

func (ctrl *TemplateController) DeleteTemplate(c *gin.Context) {
	// TODO: Implement delete template (admin only)
	c.JSON(http.StatusOK, gin.H{"message": "Delete template - TODO"})
}

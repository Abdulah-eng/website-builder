package controllers

import (
	"net/http"
	"strconv"

	"storemaker-backend/models"
	"storemaker-backend/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StoreController struct {
	db *gorm.DB
}

func NewStoreController(db *gorm.DB) *StoreController {
	return &StoreController{db: db}
}

func (ctrl *StoreController) GetUserStores(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	var stores []models.Store
	if err := ctrl.db.Where("owner_id = ?", userID).Find(&stores).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch stores"})
		return
	}

	// Return stores directly as array
	c.JSON(http.StatusOK, stores)
}

func (ctrl *StoreController) CreateStore(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	var req models.StoreCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate slug from name
	baseSlug := utils.GenerateSlug(req.Name)
	slug := utils.GenerateUniqueSlug(baseSlug, func(s string) bool {
		var count int64
		ctrl.db.Model(&models.Store{}).Where("slug = ?", s).Count(&count)
		return count > 0
	})

	// Generate subdomain
	subdomain := utils.GenerateUniqueSlug(baseSlug, func(s string) bool {
		var count int64
		ctrl.db.Model(&models.Store{}).Where("subdomain = ?", s).Count(&count)
		return count > 0
	})

	store := models.Store{
		Name:        req.Name,
		Slug:        slug,
		Subdomain:   subdomain,
		Description: req.Description,
		OwnerID:     userID.(uint),
		TemplateID:  req.TemplateID,
		Status:      models.StoreStatusDraft,
	}

	if err := ctrl.db.Create(&store).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create store"})
		return
	}

	c.JSON(http.StatusCreated, store)
}

func (ctrl *StoreController) GetStore(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	storeID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
		return
	}

	var store models.Store
	if err := ctrl.db.Where("id = ? AND owner_id = ?", storeID, userID).First(&store).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch store"})
		}
		return
	}

	c.JSON(http.StatusOK, store)
}

func (ctrl *StoreController) GetStoreBySlug(c *gin.Context) {
	slug := c.Param("slug")

	var store models.Store
	if err := ctrl.db.Where("slug = ? AND status = ?", slug, models.StoreStatusActive).First(&store).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch store"})
		}
		return
	}

	c.JSON(http.StatusOK, store)
}

func (ctrl *StoreController) UpdateStore(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	storeID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
		return
	}

	var store models.Store
	if err := ctrl.db.Where("id = ? AND owner_id = ?", storeID, userID).First(&store).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch store"})
		}
		return
	}

	var req models.StoreUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields if provided
	if req.Name != nil {
		store.Name = *req.Name
		// Update slug if name changed
		baseSlug := utils.GenerateSlug(*req.Name)
		newSlug := utils.GenerateUniqueSlug(baseSlug, func(s string) bool {
			if s == store.Slug {
				return false // Allow keeping the same slug
			}
			var count int64
			ctrl.db.Model(&models.Store{}).Where("slug = ?", s).Count(&count)
			return count > 0
		})
		store.Slug = newSlug
	}
	if req.Description != nil {
		store.Description = *req.Description
	}
	if req.Logo != nil {
		store.Logo = *req.Logo
	}
	if req.Favicon != nil {
		store.Favicon = *req.Favicon
	}
	if req.Domain != nil {
		store.Domain = *req.Domain
	}
	if req.Status != nil {
		store.Status = *req.Status
	}

	if err := ctrl.db.Save(&store).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update store"})
		return
	}

	c.JSON(http.StatusOK, store)
}

func (ctrl *StoreController) DeleteStore(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	storeID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
		return
	}

	var store models.Store
	if err := ctrl.db.Where("id = ? AND owner_id = ?", storeID, userID).First(&store).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch store"})
		}
		return
	}

	if err := ctrl.db.Delete(&store).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete store"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Store deleted successfully"})
}

func (ctrl *StoreController) GetStoreAnalytics(c *gin.Context) {
	// TODO: Implement get store analytics
	c.JSON(http.StatusOK, gin.H{"message": "Get store analytics - TODO"})
}

func (ctrl *StoreController) GetAllStores(c *gin.Context) {
	// TODO: Implement get all stores (admin only)
	c.JSON(http.StatusOK, gin.H{"message": "Get all stores - TODO"})
}

func (ctrl *StoreController) UpdateStoreStatus(c *gin.Context) {
	// TODO: Implement update store status (admin only)
	c.JSON(http.StatusOK, gin.H{"message": "Update store status - TODO"})
}

func (ctrl *StoreController) GetSystemAnalytics(c *gin.Context) {
	// TODO: Implement get system analytics (admin only)
	c.JSON(http.StatusOK, gin.H{"message": "Get system analytics - TODO"})
}

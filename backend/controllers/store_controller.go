package controllers

import (
	"encoding/json"
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

	// Generate unique domain (using subdomain with a domain suffix)
	domain := utils.GenerateUniqueSlug(baseSlug, func(s string) bool {
		var count int64
		ctrl.db.Model(&models.Store{}).Where("domain = ?", s+".storemaker.com").Count(&count)
		return count > 0
	}) + ".storemaker.com"

	store := models.Store{
		Name:        req.Name,
		Slug:        slug,
		Subdomain:   subdomain,
		Domain:      domain,
		Description: req.Description,
		OwnerID:     userID.(uint),
		TemplateID:  req.TemplateID,
		Status:      models.StoreStatusDraft,
	}

	// Save the store first
	if err := ctrl.db.Create(&store).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create store"})
		return
	}

	// If a template is selected, copy its theme, settings, and pages to the new store
	if req.TemplateID != nil {
		var template models.Template
		if err := ctrl.db.Preload("Stores").First(&template, *req.TemplateID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Template not found"})
			return
		}

		// Copy theme (if template.Config has theme info)
		if themeConfig, ok := template.Config["theme"].(map[string]interface{}); ok {
			theme := models.StoreTheme{
				StoreID: store.ID,
			}
			if colors, ok := themeConfig["colors"].(map[string]interface{}); ok {
				theme.Colors = make(models.ThemeColors)
				for k, v := range colors {
					if s, ok := v.(string); ok {
						theme.Colors[k] = s
					}
				}
			}
			if fonts, ok := themeConfig["fonts"].(map[string]interface{}); ok {
				theme.Fonts = make(models.ThemeFonts)
				for k, v := range fonts {
					if s, ok := v.(string); ok {
						theme.Fonts[k] = s
					}
				}
			}
			if v, ok := themeConfig["logo_url"].(string); ok {
				theme.LogoURL = v
			}
			if v, ok := themeConfig["favicon_url"].(string); ok {
				theme.FaviconURL = v
			}
			if v, ok := themeConfig["header_style"].(string); ok {
				theme.HeaderStyle = v
			}
			if v, ok := themeConfig["footer_style"].(string); ok {
				theme.FooterStyle = v
			}
			if v, ok := themeConfig["button_style"].(string); ok {
				theme.ButtonStyle = v
			}
			if v, ok := themeConfig["custom_css"].(string); ok {
				theme.CustomCSS = v
			}
			if err := ctrl.db.Create(&theme).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy template theme"})
				return
			}
		}

		// Copy settings (if template.Config has settings info)
		if settingsConfig, ok := template.Config["settings"].(map[string]interface{}); ok {
			settings := models.StoreSettings{
				StoreID: store.ID,
			}
			if v, ok := settingsConfig["currency"].(string); ok {
				settings.Currency = v
			}
			if v, ok := settingsConfig["language"].(string); ok {
				settings.Language = v
			}
			if v, ok := settingsConfig["timezone"].(string); ok {
				settings.Timezone = v
			}
			if v, ok := settingsConfig["allow_guest_checkout"].(bool); ok {
				settings.AllowGuestCheckout = v
			}
			if v, ok := settingsConfig["require_shipping"].(bool); ok {
				settings.RequireShipping = v
			}
			if v, ok := settingsConfig["tax_included"].(bool); ok {
				settings.TaxIncluded = v
			}
			if v, ok := settingsConfig["tax_rate"].(float64); ok {
				settings.TaxRate = v
			}
			if v, ok := settingsConfig["shipping_rate"].(float64); ok {
				settings.ShippingRate = v
			}
			if v, ok := settingsConfig["free_shipping_min"].(float64); ok {
				settings.FreeShippingMin = v
			}
			if v, ok := settingsConfig["order_prefix"].(string); ok {
				settings.OrderPrefix = v
			}
			if err := ctrl.db.Create(&settings).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy template settings"})
				return
			}
		}

		// Copy pages (if template.Config has pages info)
		if pagesConfig, ok := template.Config["pages"].([]interface{}); ok {
			for _, pageRaw := range pagesConfig {
				pageMap, ok := pageRaw.(map[string]interface{})
				if !ok {
					continue
				}
				page := models.Page{
					StoreID:     store.ID,
					Title:       getString(pageMap, "title"),
					Slug:        getString(pageMap, "slug"),
					Content:     getString(pageMap, "content"),
					Type:        models.PageType(getString(pageMap, "type")),
					IsPublished: getBool(pageMap, "is_published"),
					SeoTitle:    getString(pageMap, "seo_title"),
					SeoDesc:     getString(pageMap, "seo_description"),
				}
				if err := ctrl.db.Create(&page).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy template page"})
					return
				}
				// Copy sections
				if sectionsConfig, ok := pageMap["sections"].([]interface{}); ok {
					for _, sectionRaw := range sectionsConfig {
						sectionMap, ok := sectionRaw.(map[string]interface{})
						if !ok {
							continue
						}
						section := models.Section{
							PageID:    page.ID,
							Name:      getString(sectionMap, "name"),
							Type:      models.SectionType(getString(sectionMap, "type")),
							Order:     getInt(sectionMap, "order"),
							IsVisible: getBool(sectionMap, "is_visible"),
						}
						if err := ctrl.db.Create(&section).Error; err != nil {
							c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy template section"})
							return
						}
						// Copy components
						if componentsConfig, ok := sectionMap["components"].([]interface{}); ok {
							for _, compRaw := range componentsConfig {
								compMap, ok := compRaw.(map[string]interface{})
								if !ok {
									continue
								}
								component := models.Component{
									SectionID: section.ID,
									Name:      getString(compMap, "name"),
									Type:      models.ComponentType(getString(compMap, "type")),
									Order:     getInt(compMap, "order"),
									IsVisible: getBool(compMap, "is_visible"),
								}
								// Config is a map
								if config, ok := compMap["config"].(map[string]interface{}); ok {
									configBytes, _ := json.Marshal(config)
									var compConfig models.ComponentConfig
									_ = json.Unmarshal(configBytes, &compConfig)
									component.Config = compConfig
								}
								if err := ctrl.db.Create(&component).Error; err != nil {
									c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy template component"})
									return
								}
							}
						}
					}
				}
			}
		}
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

// Helper functions for type assertions
func getString(m map[string]interface{}, key string) string {
	if v, ok := m[key]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func getBool(m map[string]interface{}, key string) bool {
	if v, ok := m[key]; ok {
		if b, ok := v.(bool); ok {
			return b
		}
	}
	return false
}

func getInt(m map[string]interface{}, key string) int {
	if v, ok := m[key]; ok {
		switch t := v.(type) {
		case int:
			return t
		case float64:
			return int(t)
		}
	}
	return 0
}

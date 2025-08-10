package controllers

import (
	"net/http"
	"strconv"

	"storemaker-backend/models"
	"storemaker-backend/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CustomizationController struct {
	db *gorm.DB
}

func NewCustomizationController(db *gorm.DB) *CustomizationController {
	return &CustomizationController{db: db}
}

// Store Layout Management
type StoreLayoutResponse struct {
	Components []ComponentData    `json:"components"`
	Theme      *models.StoreTheme `json:"theme"`
}

type ComponentData struct {
	ID    string                 `json:"id"`
	Type  string                 `json:"type"`
	Props map[string]interface{} `json:"props"`
	Order int                    `json:"order"`
}

type StoreLayoutRequest struct {
	Components []ComponentData    `json:"components"`
	Theme      *models.StoreTheme `json:"theme"`
}

func (ctrl *CustomizationController) GetStoreLayout(c *gin.Context) {
	storeID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
		return
	}

	// Check if store exists
	var store models.Store
	if err := ctrl.db.First(&store, storeID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch store"})
		}
		return
	}

	// Get home page for this store
	var homePage models.Page
	if err := ctrl.db.Where("store_id = ? AND type = ?", storeID, models.PageTypeHome).
		Preload("Sections.Components").First(&homePage).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Check if there are any pages for this store (might be from template)
			var anyPage models.Page
			if err := ctrl.db.Where("store_id = ?", storeID).
				Preload("Sections.Components").First(&anyPage).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					// No pages exist, create default home page
					homePage = models.Page{
						Title:       "Home",
						Slug:        "home",
						Type:        models.PageTypeHome,
						StoreID:     uint(storeID),
						IsPublished: true,
					}
					if err := ctrl.db.Create(&homePage).Error; err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create home page"})
						return
					}
				} else {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch store layout"})
					return
				}
			} else {
				// Use the first page found (likely from template)
				homePage = anyPage
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch store layout"})
			return
		}
	}

	// Get store theme
	var theme models.StoreTheme
	if err := ctrl.db.Where("store_id = ?", storeID).First(&theme).Error; err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch store theme"})
		return
	}

	// Convert sections/components to frontend format
	var components []ComponentData
	for _, section := range homePage.Sections {
		for _, component := range section.Components {
			components = append(components, ComponentData{
				ID:    component.Name, // Use name as ID for now
				Type:  string(component.Type),
				Props: component.Config,
				Order: component.Order,
			})
		}
	}

	response := StoreLayoutResponse{
		Components: components,
		Theme:      &theme,
	}

	c.JSON(http.StatusOK, response)
}

// Public endpoint - Get store layout by slug (for public store access)
func (ctrl *CustomizationController) GetPublicStoreLayout(c *gin.Context) {
	storeSlug := c.Param("slug")

	// Get store by slug
	var store models.Store
	if err := ctrl.db.Where("slug = ? AND status = ?", storeSlug, models.StoreStatusActive).First(&store).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch store"})
		}
		return
	}

	// Get home page for this store
	var homePage models.Page
	if err := ctrl.db.Where("store_id = ? AND type = ?", store.ID, models.PageTypeHome).
		Preload("Sections.Components").First(&homePage).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Check if there are any pages for this store (might be from template)
			var anyPage models.Page
			if err := ctrl.db.Where("store_id = ?", store.ID).
				Preload("Sections.Components").First(&anyPage).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					// No pages exist, return empty layout
					response := StoreLayoutResponse{
						Components: []ComponentData{},
						Theme:      nil,
					}
					c.JSON(http.StatusOK, response)
					return
				} else {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch store layout"})
					return
				}
			} else {
				// Use the first page found (likely from template)
				homePage = anyPage
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch store layout"})
			return
		}
	}

	// Get store theme
	var theme models.StoreTheme
	if err := ctrl.db.Where("store_id = ?", store.ID).First(&theme).Error; err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch store theme"})
		return
	}

	// Convert sections/components to frontend format
	var components []ComponentData
	for _, section := range homePage.Sections {
		for _, component := range section.Components {
			components = append(components, ComponentData{
				ID:    component.Name, // Use name as ID for now
				Type:  string(component.Type),
				Props: component.Config,
				Order: component.Order,
			})
		}
	}

	response := StoreLayoutResponse{
		Components: components,
		Theme:      &theme,
	}

	c.JSON(http.StatusOK, response)
}

func (ctrl *CustomizationController) SaveStoreLayout(c *gin.Context) {
	storeID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
		return
	}

	var req StoreLayoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Start transaction
	tx := ctrl.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Get or create home page
	var homePage models.Page
	if err := tx.Where("store_id = ? AND type = ?", storeID, models.PageTypeHome).First(&homePage).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			homePage = models.Page{
				Title:       "Home",
				Slug:        "home",
				Type:        models.PageTypeHome,
				StoreID:     uint(storeID),
				IsPublished: true,
			}
			if err := tx.Create(&homePage).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create home page"})
				return
			}
		} else {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch home page"})
			return
		}
	}

	// Clear existing sections and components
	if err := tx.Where("page_id = ?", homePage.ID).Delete(&models.Section{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear existing layout"})
		return
	}

	// Create new section for components
	section := models.Section{
		Name:      "Main Content",
		Type:      models.SectionTypeText, // Generic type
		PageID:    homePage.ID,
		Order:     0,
		IsVisible: true,
	}
	if err := tx.Create(&section).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create section"})
		return
	}

	// Create components
	for _, comp := range req.Components {
		component := models.Component{
			Name:      comp.ID,
			Type:      models.ComponentType(comp.Type),
			Config:    models.ComponentConfig(comp.Props),
			Order:     comp.Order,
			SectionID: section.ID,
			IsVisible: true,
		}
		if err := tx.Create(&component).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create component"})
			return
		}
	}

	// Update theme if provided
	if req.Theme != nil {
		var existingTheme models.StoreTheme
		if err := tx.Where("store_id = ?", storeID).First(&existingTheme).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Create new theme
				req.Theme.StoreID = uint(storeID)
				if err := tx.Create(req.Theme).Error; err != nil {
					tx.Rollback()
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create theme"})
					return
				}
			} else {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch existing theme"})
				return
			}
		} else {
			// Update existing theme
			req.Theme.ID = existingTheme.ID
			req.Theme.StoreID = uint(storeID)
			if err := tx.Save(req.Theme).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update theme"})
				return
			}
		}
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save layout"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Layout saved successfully"})
}

// Public endpoint - Get published pages by store slug
func (ctrl *CustomizationController) GetStorePages(c *gin.Context) {
	storeSlug := c.Param("slug")

	var store models.Store
	if err := ctrl.db.Where("slug = ?", storeSlug).First(&store).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		return
	}

	var pages []models.Page
	if err := ctrl.db.Where("store_id = ? AND is_published = ?", store.ID, true).Find(&pages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch pages"})
		return
	}

	c.JSON(http.StatusOK, pages)
}

// Management endpoint - Get all pages for store owner
func (ctrl *CustomizationController) GetPages(c *gin.Context) {
	storeID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
		return
	}

	var pages []models.Page
	query := ctrl.db.Where("store_id = ?", storeID)

	// Filter by type if provided
	if pageType := c.Query("type"); pageType != "" {
		query = query.Where("type = ?", pageType)
	}

	// Filter by published status if provided
	if published := c.Query("published"); published != "" {
		switch published {
		case "true":
			query = query.Where("is_published = ?", true)
		case "false":
			query = query.Where("is_published = ?", false)
		}
	}

	if err := query.Order("created_at DESC").Find(&pages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch pages"})
		return
	}

	c.JSON(http.StatusOK, pages)
}

func (ctrl *CustomizationController) GetStorePage(c *gin.Context) {
	storeSlug := c.Param("slug")
	pageSlug := c.Param("pageSlug")

	var store models.Store
	if err := ctrl.db.Where("slug = ?", storeSlug).First(&store).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		return
	}

	var page models.Page
	if err := ctrl.db.Where("store_id = ? AND slug = ? AND is_published = ?", store.ID, pageSlug, true).
		Preload("Sections.Components").First(&page).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
		return
	}

	c.JSON(http.StatusOK, page)
}

func (ctrl *CustomizationController) GetPublicStorePages(c *gin.Context) {
	storeSlug := c.Param("slug")

	var store models.Store
	if err := ctrl.db.Where("slug = ?", storeSlug).First(&store).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		return
	}

	var pages []models.Page
	if err := ctrl.db.Where("store_id = ? AND is_published = ?", store.ID, true).
		Preload("Sections.Components").
		Find(&pages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch pages"})
		return
	}

	c.JSON(http.StatusOK, pages)
}

func (ctrl *CustomizationController) GetPublicStorePage(c *gin.Context) {
	storeSlug := c.Param("slug")
	pageSlug := c.Param("pageSlug")

	var store models.Store
	if err := ctrl.db.Where("slug = ?", storeSlug).First(&store).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		return
	}

	var page models.Page
	if err := ctrl.db.Where("store_id = ? AND slug = ? AND is_published = ?", store.ID, pageSlug, true).
		Preload("Sections.Components").
		First(&page).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
		return
	}

	c.JSON(http.StatusOK, page)
}

func (ctrl *CustomizationController) CreatePage(c *gin.Context) {
	storeID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
		return
	}

	var req models.Page
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate slug from title if not provided
	if req.Slug == "" && req.Title != "" {
		baseSlug := utils.GenerateSlug(req.Title)
		req.Slug = utils.GenerateUniqueSlug(baseSlug, func(s string) bool {
			var count int64
			ctrl.db.Model(&models.Page{}).Where("store_id = ? AND slug = ?", storeID, s).Count(&count)
			return count > 0
		})
	}

	req.StoreID = uint(storeID)
	if err := ctrl.db.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create page"})
		return
	}

	c.JSON(http.StatusCreated, req)
}

func (ctrl *CustomizationController) UpdatePage(c *gin.Context) {
	pageID, err := strconv.ParseUint(c.Param("pageId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page ID"})
		return
	}

	var req models.Page
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.db.Model(&models.Page{}).Where("id = ?", pageID).Updates(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update page"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Page updated successfully"})
}

func (ctrl *CustomizationController) DeletePage(c *gin.Context) {
	pageID, err := strconv.ParseUint(c.Param("pageId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page ID"})
		return
	}

	if err := ctrl.db.Delete(&models.Page{}, pageID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete page"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Page deleted successfully"})
}

func (ctrl *CustomizationController) GetStoreSettings(c *gin.Context) {
	storeID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
		return
	}

	var settings models.StoreSettings
	if err := ctrl.db.Where("store_id = ?", storeID).First(&settings).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create default settings
			settings = models.StoreSettings{
				StoreID:            uint(storeID),
				Currency:           "USD",
				Language:           "en",
				Timezone:           "UTC",
				AllowGuestCheckout: true,
				RequireShipping:    true,
			}
			if err := ctrl.db.Create(&settings).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create default settings"})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch settings"})
			return
		}
	}

	c.JSON(http.StatusOK, settings)
}

func (ctrl *CustomizationController) UpdateStoreSettings(c *gin.Context) {
	storeID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
		return
	}

	var req models.StoreSettings
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.db.Model(&models.StoreSettings{}).Where("store_id = ?", storeID).Updates(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update settings"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Settings updated successfully"})
}

func (ctrl *CustomizationController) GetStoreTheme(c *gin.Context) {
	storeID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
		return
	}

	var theme models.StoreTheme
	if err := ctrl.db.Where("store_id = ?", storeID).First(&theme).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create default theme
			theme = models.StoreTheme{
				StoreID: uint(storeID),
				Colors:  make(models.ThemeColors),
				Fonts:   make(models.ThemeFonts),
			}
			if err := ctrl.db.Create(&theme).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create default theme"})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch theme"})
			return
		}
	}

	c.JSON(http.StatusOK, theme)
}

func (ctrl *CustomizationController) UpdateStoreTheme(c *gin.Context) {
	storeID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
		return
	}

	var req models.StoreTheme
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.db.Model(&models.StoreTheme{}).Where("store_id = ?", storeID).Updates(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update theme"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Theme updated successfully"})
}

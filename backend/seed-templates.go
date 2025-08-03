package main

import (
	"log"

	"storemaker-backend/config"
	"storemaker-backend/database"
	"storemaker-backend/models"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db, err := database.Initialize(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Sample templates
	templates := []models.Template{
		{
			Name:        "Fashion Store",
			Description: "Modern and elegant template perfect for fashion brands",
			Category:    "fashion",
			Config: models.TemplateConfig{
				"colors": map[string]interface{}{
					"primary": "#1f2937",
					"secondary": "#f3f4f6",
				},
			},
			IsActive:  true,
			IsPremium: false,
		},
		{
			Name:        "Electronics Hub",
			Description: "Clean tech-focused template for electronics stores",
			Category:    "electronics",
			Config: models.TemplateConfig{
				"colors": map[string]interface{}{
					"primary": "#3b82f6",
					"secondary": "#e5e7eb",
				},
			},
			IsActive:  true,
			IsPremium: false,
		},
		{
			Name:        "Food & Restaurant",
			Description: "Appetizing template for restaurants and food delivery",
			Category:    "food",
			Config: models.TemplateConfig{
				"colors": map[string]interface{}{
					"primary": "#dc2626",
					"secondary": "#fef2f2",
				},
			},
			IsActive:  true,
			IsPremium: false,
		},
	}

	// Insert templates
	for _, template := range templates {
		var existing models.Template
		result := db.Where("name = ?", template.Name).First(&existing)
		if result.Error != nil {
			// Template doesn't exist, create it
			if err := db.Create(&template).Error; err != nil {
				log.Printf("Failed to create template %s: %v", template.Name, err)
			} else {
				log.Printf("Created template: %s", template.Name)
			}
		} else {
			log.Printf("Template already exists: %s", template.Name)
		}
	}

	log.Println("Template seeding completed!")
}
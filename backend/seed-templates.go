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

	// Sample templates with comprehensive configurations
	templates := []models.Template{
		{
			Name:        "Fashion Store",
			Description: "Modern and elegant template perfect for fashion brands",
			Category:    "fashion",
			Config: models.TemplateConfig{
				"theme": map[string]interface{}{
					"colors": map[string]interface{}{
						"primary": "#1f2937",
						"secondary": "#f3f4f6",
						"accent": "#f59e0b",
						"text": "#1f2937",
						"background": "#ffffff",
					},
					"fonts": map[string]interface{}{
						"heading": "Playfair Display",
						"body": "Inter",
					},
					"logo_url": "",
					"favicon_url": "",
					"header_style": "modern",
					"footer_style": "detailed",
					"button_style": "rounded",
					"custom_css": "",
				},
				"settings": map[string]interface{}{
					"currency": "USD",
					"language": "en",
					"timezone": "UTC",
					"allow_guest_checkout": true,
					"require_shipping": true,
					"tax_included": false,
					"tax_rate": 8.5,
					"shipping_rate": 5.99,
					"free_shipping_min": 50.0,
					"order_prefix": "#",
				},
				"pages": []interface{}{
					map[string]interface{}{
						"title": "Home",
						"slug": "home",
						"content": "",
						"type": "home",
						"is_published": true,
						"seo_title": "Fashion Store - Home",
						"seo_description": "Discover the latest fashion trends and styles",
						"sections": []interface{}{
							map[string]interface{}{
								"name": "Hero Section",
								"type": "hero",
								"order": 0,
								"is_visible": true,
								"components": []interface{}{
									map[string]interface{}{
										"name": "hero-banner",
										"type": "hero-banner",
										"order": 0,
										"is_visible": true,
										"config": map[string]interface{}{
											"title": "Elegant Fashion Collection",
											"subtitle": "Discover timeless styles for every occasion",
											"buttonText": "Shop Collection",
											"backgroundImage": "https://images.unsplash.com/photo-1441986300917-64674bd600d8?w=1200&h=600&fit=crop",
											"overlay": true,
											"secondaryButton": true,
											"secondaryButtonText": "Learn More",
										},
									},
								},
							},
							map[string]interface{}{
								"name": "Featured Products",
								"type": "products",
								"order": 1,
								"is_visible": true,
								"components": []interface{}{
									map[string]interface{}{
										"name": "product-grid",
										"type": "product-grid",
										"order": 0,
										"is_visible": true,
										"config": map[string]interface{}{
											"title": "Featured Products",
											"columns": 3,
											"showPrice": true,
											"showRating": true,
											"maxProducts": 6,
										},
									},
								},
							},
							map[string]interface{}{
								"name": "Features",
								"type": "features",
								"order": 2,
								"is_visible": true,
								"components": []interface{}{
									map[string]interface{}{
										"name": "feature-list",
										"type": "feature-list",
										"order": 0,
										"is_visible": true,
										"config": map[string]interface{}{
											"title": "Why Choose Us",
											"features": []interface{}{
												map[string]interface{}{
													"icon": "shipping",
													"title": "Free Shipping",
													"description": "On orders over $100",
												},
												map[string]interface{}{
													"icon": "support",
													"title": "24/7 Support",
													"description": "Always here to help",
												},
												map[string]interface{}{
													"icon": "returns",
													"title": "Easy Returns",
													"description": "30-day return policy",
												},
											},
										},
									},
								},
							},
						},
					},
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
				"theme": map[string]interface{}{
					"colors": map[string]interface{}{
						"primary": "#3b82f6",
						"secondary": "#e5e7eb",
						"accent": "#10b981",
						"text": "#1f2937",
						"background": "#ffffff",
					},
					"fonts": map[string]interface{}{
						"heading": "Inter",
						"body": "Inter",
					},
					"logo_url": "",
					"favicon_url": "",
					"header_style": "modern",
					"footer_style": "detailed",
					"button_style": "rounded",
					"custom_css": "",
				},
				"settings": map[string]interface{}{
					"currency": "USD",
					"language": "en",
					"timezone": "UTC",
					"allow_guest_checkout": true,
					"require_shipping": true,
					"tax_included": false,
					"tax_rate": 7.5,
					"shipping_rate": 8.99,
					"free_shipping_min": 75.0,
					"order_prefix": "#",
				},
				"pages": []interface{}{
					map[string]interface{}{
						"title": "Home",
						"slug": "home",
						"content": "",
						"type": "home",
						"is_published": true,
						"seo_title": "Electronics Hub - Home",
						"seo_description": "Latest electronics and tech gadgets",
						"sections": []interface{}{
							map[string]interface{}{
								"name": "Hero Section",
								"type": "hero",
								"order": 0,
								"is_visible": true,
								"components": []interface{}{
									map[string]interface{}{
										"name": "hero-split",
										"type": "hero-split",
										"order": 0,
										"is_visible": true,
										"config": map[string]interface{}{
											"title": "Latest Tech Gadgets",
											"subtitle": "Discover cutting-edge electronics",
											"buttonText": "Shop Now",
											"image": "https://images.unsplash.com/photo-1560472354-b33ff0c44a43?w=800&h=600&fit=crop",
											"imagePosition": "right",
										},
									},
								},
							},
							map[string]interface{}{
								"name": "Product Categories",
								"type": "products",
								"order": 1,
								"is_visible": true,
								"components": []interface{}{
									map[string]interface{}{
										"name": "product-categories",
										"type": "product-categories",
										"order": 0,
										"is_visible": true,
										"config": map[string]interface{}{
											"title": "Shop by Category",
											"categories": []interface{}{
												map[string]interface{}{
													"name": "Smartphones",
													"image": "",
													"itemCount": 25,
												},
												map[string]interface{}{
													"name": "Laptops",
													"image": "",
													"itemCount": 18,
												},
												map[string]interface{}{
													"name": "Accessories",
													"image": "",
													"itemCount": 45,
												},
												map[string]interface{}{
													"name": "Gaming",
													"image": "",
													"itemCount": 32,
												},
											},
										},
									},
								},
							},
							map[string]interface{}{
								"name": "Featured Products",
								"type": "products",
								"order": 2,
								"is_visible": true,
								"components": []interface{}{
									map[string]interface{}{
										"name": "product-carousel",
										"type": "product-carousel",
										"order": 0,
										"is_visible": true,
										"config": map[string]interface{}{
											"title": "Trending Now",
											"autoplay": true,
											"showDots": true,
											"slidesToShow": 4,
										},
									},
								},
							},
						},
					},
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
				"theme": map[string]interface{}{
					"colors": map[string]interface{}{
						"primary": "#dc2626",
						"secondary": "#fef2f2",
						"accent": "#f59e0b",
						"text": "#1f2937",
						"background": "#ffffff",
					},
					"fonts": map[string]interface{}{
						"heading": "Playfair Display",
						"body": "Inter",
					},
					"logo_url": "",
					"favicon_url": "",
					"header_style": "modern",
					"footer_style": "detailed",
					"button_style": "rounded",
					"custom_css": "",
				},
				"settings": map[string]interface{}{
					"currency": "USD",
					"language": "en",
					"timezone": "UTC",
					"allow_guest_checkout": true,
					"require_shipping": false,
					"tax_included": true,
					"tax_rate": 0.0,
					"shipping_rate": 0.0,
					"free_shipping_min": 0.0,
					"order_prefix": "#",
				},
				"pages": []interface{}{
					map[string]interface{}{
						"title": "Home",
						"slug": "home",
						"content": "",
						"type": "home",
						"is_published": true,
						"seo_title": "Food & Restaurant - Home",
						"seo_description": "Delicious food and exceptional dining experience",
						"sections": []interface{}{
							map[string]interface{}{
								"name": "Hero Section",
								"type": "hero",
								"order": 0,
								"is_visible": true,
								"components": []interface{}{
									map[string]interface{}{
										"name": "hero-banner",
										"type": "hero-banner",
										"order": 0,
										"is_visible": true,
										"config": map[string]interface{}{
											"title": "Delicious Food Delivered",
											"subtitle": "Fresh ingredients, amazing taste",
											"buttonText": "Order Now",
											"backgroundImage": "https://images.unsplash.com/photo-1565299624946-b28f40a0ca4b?w=1200&h=600&fit=crop",
											"overlay": true,
											"secondaryButton": false,
											"secondaryButtonText": "",
										},
									},
								},
							},
							map[string]interface{}{
								"name": "About Section",
								"type": "info",
								"order": 1,
								"is_visible": true,
								"components": []interface{}{
									map[string]interface{}{
										"name": "about-section",
										"type": "about-section",
										"order": 0,
										"is_visible": true,
										"config": map[string]interface{}{
											"title": "Our Story",
											"content": "We are passionate about creating delicious meals using the finest ingredients. Our chefs bring years of experience and creativity to every dish.",
											"image": "",
											"imagePosition": "left",
										},
									},
								},
							},
							map[string]interface{}{
								"name": "Contact Info",
								"type": "info",
								"order": 2,
								"is_visible": true,
								"components": []interface{}{
									map[string]interface{}{
										"name": "contact-info",
										"type": "contact-info",
										"order": 0,
										"is_visible": true,
										"config": map[string]interface{}{
											"title": "Visit Our Restaurant",
											"address": "123 Food Street, City, State 12345",
											"phone": "(555) 123-4567",
											"email": "info@restaurant.com",
											"hours": "Mon-Sun: 11AM-10PM",
										},
									},
								},
							},
						},
					},
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
			// Update existing template with new configuration
			if err := db.Model(&existing).Updates(map[string]interface{}{
				"description": template.Description,
				"category":    template.Category,
				"config":      template.Config,
				"is_active":   template.IsActive,
				"is_premium":  template.IsPremium,
			}).Error; err != nil {
				log.Printf("Failed to update template %s: %v", template.Name, err)
			} else {
				log.Printf("Updated template: %s", template.Name)
			}
		}
	}

	log.Println("Template seeding completed!")
}
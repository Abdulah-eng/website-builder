package routes

import (
	"storemaker-backend/controllers"
	"storemaker-backend/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// Initialize controllers
	authController := controllers.NewAuthController(db)
	userController := controllers.NewUserController(db)
	storeController := controllers.NewStoreController(db)
	templateController := controllers.NewTemplateController(db)
	productController := controllers.NewProductController(db)
	orderController := controllers.NewOrderController(db)
	customizationController := controllers.NewCustomizationController(db)

	// API v1 routes
	v1 := router.Group("/api/v1")

	// Public routes
	public := v1.Group("")
	{
		// Authentication
		public.POST("/auth/register", authController.Register)
		public.POST("/auth/login", authController.Login)
		public.POST("/auth/refresh", authController.RefreshToken)

		// Public store access
		public.GET("/stores/:slug", storeController.GetStoreBySlug)
		public.GET("/stores/:slug/layout", customizationController.GetPublicStoreLayout)
		public.GET("/stores/:slug/products", productController.GetStoreProducts)
		public.GET("/stores/:slug/products/:productSlug", productController.GetStoreProduct)
		public.GET("/stores/:slug/pages", customizationController.GetStorePages)
		public.GET("/stores/:slug/pages/:pageSlug", customizationController.GetStorePage)

		// Public templates
		public.GET("/templates", templateController.GetPublicTemplates)
		public.GET("/templates/:id", templateController.GetTemplate)

		// Public orders (for guests)
		public.POST("/stores/:slug/orders", orderController.CreateOrder)
		public.GET("/orders/:orderNumber", orderController.GetOrderByNumber)
	}

	// Protected routes (require authentication)
	protected := v1.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		// User routes
		protected.GET("/user/profile", userController.GetProfile)
		protected.PUT("/user/profile", userController.UpdateProfile)
		protected.DELETE("/user/profile", userController.DeleteProfile)

		// Store management (merchant only) - using /manage prefix to avoid conflicts
		storeRoutes := protected.Group("/manage/stores")
		storeRoutes.Use(middleware.MerchantMiddleware())
		{
			storeRoutes.GET("", storeController.GetUserStores)
			storeRoutes.POST("", storeController.CreateStore)
			storeRoutes.GET("/:id", storeController.GetStore)
			storeRoutes.PUT("/:id", storeController.UpdateStore)
			storeRoutes.DELETE("/:id", storeController.DeleteStore)

			// Store settings
			storeRoutes.GET("/:id/settings", customizationController.GetStoreSettings)
			storeRoutes.PUT("/:id/settings", customizationController.UpdateStoreSettings)

			// Store theme
			storeRoutes.GET("/:id/theme", customizationController.GetStoreTheme)
			storeRoutes.PUT("/:id/theme", customizationController.UpdateStoreTheme)

			// Store layout
			storeRoutes.GET("/:id/layout", customizationController.GetStoreLayout)
			storeRoutes.POST("/:id/layout", customizationController.SaveStoreLayout)

			// Store pages
			storeRoutes.GET("/:id/pages", customizationController.GetPages)
			storeRoutes.POST("/:id/pages", customizationController.CreatePage)
			storeRoutes.PUT("/:id/pages/:pageId", customizationController.UpdatePage)
			storeRoutes.DELETE("/:id/pages/:pageId", customizationController.DeletePage)

			// Store products
			storeRoutes.GET("/:id/products", productController.GetProducts)
			storeRoutes.POST("/:id/products", productController.CreateProduct)
			storeRoutes.PUT("/:id/products/:productId", productController.UpdateProduct)
			storeRoutes.DELETE("/:id/products/:productId", productController.DeleteProduct)

			// Store orders
			storeRoutes.GET("/:id/orders", orderController.GetStoreOrders)
			storeRoutes.PUT("/:id/orders/:orderId", orderController.UpdateOrder)

			// Store analytics
			storeRoutes.GET("/:id/analytics", storeController.GetStoreAnalytics)
		}
	}

	// Admin routes (admin only)
	admin := v1.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		// User management
		admin.GET("/users", userController.GetAllUsers)
		admin.PUT("/users/:id/status", userController.UpdateUserStatus)

		// Template management
		admin.GET("/templates", templateController.GetAllTemplates)
		admin.POST("/templates", templateController.CreateTemplate)
		admin.PUT("/templates/:id", templateController.UpdateTemplate)
		admin.DELETE("/templates/:id", templateController.DeleteTemplate)

		// Store management (admin view of all stores)
		admin.GET("/stores", storeController.GetAllStores)
		admin.PUT("/stores/:id/status", storeController.UpdateStoreStatus)

		// System analytics
		admin.GET("/analytics", storeController.GetSystemAnalytics)
	}
}

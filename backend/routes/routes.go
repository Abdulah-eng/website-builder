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
	newsletterController := controllers.NewNewsletterController(db)store-builder.tsx:2037 Loaded layout data: {components: Array(2), theme: {…}}
	store-builder.tsx:2053 Loaded 2 components
	store-builder.tsx:2064 Loaded theme configuration
	component-library.tsx:496 Encountered two children with the same key, `product-showcase`. Keys should be unique so that components maintain their identity across updates. Non-unique keys may cause children to be duplicated and/or omitted — the behavior is unsupported and could change in a future version.
	overrideMethod @ hook.js:608
	error @ intercept-console-error.ts:44
	(anonymous) @ react-dom-client.development.js:5705
	runWithFiberInDEV @ react-dom-client.development.js:872
	warnOnInvalidKey @ react-dom-client.development.js:5704
	reconcileChildrenArray @ react-dom-client.development.js:5746
	reconcileChildFibersImpl @ react-dom-client.development.js:6094
	(anonymous) @ react-dom-client.development.js:6199
	reconcileChildren @ react-dom-client.development.js:8754
	beginWork @ react-dom-client.development.js:10950
	runWithFiberInDEV @ react-dom-client.development.js:872
	performUnitOfWork @ react-dom-client.development.js:15677
	workLoopSync @ react-dom-client.development.js:15497
	renderRootSync @ react-dom-client.development.js:15477
	performWorkOnRoot @ react-dom-client.development.js:14941
	performWorkOnRootViaSchedulerTask @ react-dom-client.development.js:16766
	performWorkUntilDeadline @ scheduler.development.js:45
	<div>
	exports.jsxDEV @ react-jsx-dev-runtime.development.js:345
	(anonymous) @ component-library.tsx:496
	ComponentLibrary @ component-library.tsx:495
	react_stack_bottom_frame @ react-dom-client.development.js:23552
	renderWithHooksAgain @ react-dom-client.development.js:6863
	renderWithHooks @ react-dom-client.development.js:6775
	updateFunctionComponent @ react-dom-client.development.js:9069
	beginWork @ react-dom-client.development.js:10679
	runWithFiberInDEV @ react-dom-client.development.js:872
	performUnitOfWork @ react-dom-client.development.js:15677
	workLoopSync @ react-dom-client.development.js:15497
	renderRootSync @ react-dom-client.development.js:15477
	performWorkOnRoot @ react-dom-client.development.js:14941
	performWorkOnRootViaSchedulerTask @ react-dom-client.development.js:16766
	performWorkUntilDeadline @ scheduler.development.js:45
	<ComponentLibrary>
	exports.jsxDEV @ react-jsx-dev-runtime.development.js:345
	StoreBuilder @ store-builder.tsx:2433
	react_stack_bottom_frame @ react-dom-client.development.js:23552
	renderWithHooksAgain @ react-dom-client.development.js:6863
	renderWithHooks @ react-dom-client.development.js:6775
	updateFunctionComponent @ react-dom-client.development.js:9069
	beginWork @ react-dom-client.development.js:10679
	runWithFiberInDEV @ react-dom-client.development.js:872
	performUnitOfWork @ react-dom-client.development.js:15677
	workLoopSync @ react-dom-client.development.js:15497
	renderRootSync @ react-dom-client.development.js:15477
	performWorkOnRoot @ react-dom-client.development.js:14941
	performWorkOnRootViaSchedulerTask @ react-dom-client.development.js:16766
	performWorkUntilDeadline @ scheduler.development.js:45
	<StoreBuilder>
	exports.jsxDEV @ react-jsx-dev-runtime.development.js:345
	StoreBuilderPage @ page.tsx:75
	react_stack_bottom_frame @ react-dom-client.development.js:23552
	renderWithHooksAgain @ react-dom-client.development.js:6863
	renderWithHooks @ react-dom-client.development.js:6775
	updateFunctionComponent @ react-dom-client.development.js:9069
	beginWork @ react-dom-client.development.js:10679
	runWithFiberInDEV @ react-dom-client.development.js:872
	performUnitOfWork @ react-dom-client.development.js:15677
	workLoopSync @ react-dom-client.development.js:15497
	renderRootSync @ react-dom-client.development.js:15477
	performWorkOnRoot @ react-dom-client.development.js:14941
	performWorkOnRootViaSchedulerTask @ react-dom-client.development.js:16766
	performWorkUntilDeadline @ scheduler.development.js:45
	<StoreBuilderPage>
	exports.jsx @ react-jsx-runtime.development.js:338
	ClientPageRoot @ client-page.tsx:60
	react_stack_bottom_frame @ react-dom-client.development.js:23552
	renderWithHooksAgain @ react-dom-client.development.js:6863
	renderWithHooks @ react-dom-client.development.js:6775
	updateFunctionComponent @ react-dom-client.development.js:9069
	beginWork @ react-dom-client.development.js:10628
	runWithFiberInDEV @ react-dom-client.development.js:872
	performUnitOfWork @ react-dom-client.development.js:15677
	workLoopConcurrentByScheduler @ react-dom-client.development.js:15671
	renderRootConcurrent @ react-dom-client.development.js:15646
	performWorkOnRoot @ react-dom-client.development.js:14940
	performWorkOnRootViaSchedulerTask @ react-dom-client.development.js:16766
	performWorkUntilDeadline @ scheduler.development.js:45
	"use client"
	initializeElement @ react-server-dom-turbopack-client.browser.development.js:1199
	wakeChunk @ react-server-dom-turbopack-client.browser.development.js:969
	fulfillReference @ react-server-dom-turbopack-client.browser.development.js:1303
	wakeChunk @ react-server-dom-turbopack-client.browser.development.js:970
	wakeChunkIfInitialized @ react-server-dom-turbopack-client.browser.development.js:1005
	resolveModuleChunk @ react-server-dom-turbopack-client.browser.development.js:1090
	(anonymous) @ react-server-dom-turbopack-client.browser.development.js:1937
	"use server"
	ResponseInstance @ react-server-dom-turbopack-client.browser.development.js:1863
	createResponseFromOptions @ react-server-dom-turbopack-client.browser.development.js:2851
	exports.createFromReadableStream @ react-server-dom-turbopack-client.browser.development.js:3213
	[project]/node_modules/next/dist/client/app-index.js [app-client] (ecmascript) @ app-index.tsx:157
	(anonymous) @ dev-base.ts:201
	runModuleExecutionHooks @ dev-base.ts:256
	instantiateModule @ dev-base.ts:199
	getOrInstantiateModuleFromParent @ dev-base.ts:126
	commonJsRequire @ runtime-utils.ts:291
	(anonymous) @ app-next-turbopack.ts:11
	(anonymous) @ app-bootstrap.ts:78
	loadScriptsInSequence @ app-bootstrap.ts:20
	appBootstrap @ app-bootstrap.ts:60
	[project]/node_modules/next/dist/client/app-next-turbopack.js [app-client] (ecmascript) @ app-next-turbopack.ts:10
	(anonymous) @ dev-base.ts:201
	runModuleExecutionHooks @ dev-base.ts:256
	instantiateModule @ dev-base.ts:199
	getOrInstantiateRuntimeModule @ dev-base.ts:96
	registerChunk @ runtime-backend-dom.ts:88
	await in registerChunk
	registerChunk @ runtime-base.ts:377
	(anonymous) @ dev-backend-dom.ts:126
	(anonymous) @ dev-backend-dom.ts:126
	component-library.tsx:496 Encountered two children with the same key, `product-showcase`. Keys should be unique so that components maintain their identity across updates. Non-unique keys may cause children to be duplicated and/or omitted — the behavior is unsupported and could change in a future version.
	overrideMethod @ hook.js:608
	error @ intercept-console-error.ts:44
	(anonymous) @ react-dom-client.development.js:5705
	runWithFiberInDEV @ react-dom-client.development.js:872
	warnOnInvalidKey @ react-dom-client.development.js:5704
	reconcileChildrenArray @ react-dom-client.development.js:5804
	reconcileChildFibersImpl @ react-dom-client.development.js:6094
	(anonymous) @ react-dom-client.development.js:6199
	reconcileChildren @ react-dom-client.development.js:8754
	beginWork @ react-dom-client.development.js:10950
	runWithFiberInDEV @ react-dom-client.development.js:872
	performUnitOfWork @ react-dom-client.development.js:15677
	workLoopSync @ react-dom-client.development.js:15497
	renderRootSync @ react-dom-client.development.js:15477
	performWorkOnRoot @ react-dom-client.development.js:14941
	performSyncWorkOnRoot @ react-dom-client.development.js:16781
	flushSyncWorkAcrossRoots_impl @ react-dom-client.development.js:16627
	processRootScheduleInMicrotask @ react-dom-client.development.js:16665
	(anonymous) @ react-dom-client.development.js:16800
	<div>
	exports.jsxDEV @ react-jsx-dev-runtime.development.js:345
	(anonymous) @ component-library.tsx:496
	ComponentLibrary @ component-library.tsx:495
	react_stack_bottom_frame @ react-dom-client.development.js:23552
	renderWithHooksAgain @ react-dom-client.development.js:6863
	renderWithHooks @ react-dom-client.development.js:6775
	updateFunctionComponent @ react-dom-client.development.js:9069
	beginWork @ react-dom-client.development.js:10679
	runWithFiberInDEV @ react-dom-client.development.js:872
	performUnitOfWork @ react-dom-client.development.js:15677
	workLoopSync @ react-dom-client.development.js:15497
	renderRootSync @ react-dom-client.development.js:15477
	performWorkOnRoot @ react-dom-client.development.js:14941
	performSyncWorkOnRoot @ react-dom-client.development.js:16781
	flushSyncWorkAcrossRoots_impl @ react-dom-client.development.js:16627
	processRootScheduleInMicrotask @ react-dom-client.development.js:16665
	(anonymous) @ react-dom-client.development.js:16800
	<ComponentLibrary>
	exports.jsxDEV @ react-jsx-dev-runtime.development.js:345
	StoreBuilder @ store-builder.tsx:2433
	react_stack_bottom_frame @ react-dom-client.development.js:23552
	renderWithHooksAgain @ react-dom-client.development.js:6863
	renderWithHooks @ react-dom-client.development.js:6775
	updateFunctionComponent @ react-dom-client.development.js:9069
	beginWork @ react-dom-client.development.js:10679
	runWithFiberInDEV @ react-dom-client.development.js:872
	performUnitOfWork @ react-dom-client.development.js:15677
	workLoopSync @ react-dom-client.development.js:15497
	renderRootSync @ react-dom-client.development.js:15477
	performWorkOnRoot @ react-dom-client.development.js:14941
	performWorkOnRootViaSchedulerTask @ react-dom-client.development.js:16766
	performWorkUntilDeadline @ scheduler.development.js:45
	<StoreBuilder>
	exports.jsxDEV @ react-jsx-dev-runtime.development.js:345
	StoreBuilderPage @ page.tsx:75
	react_stack_bottom_frame @ react-dom-client.development.js:23552
	renderWithHooksAgain @ react-dom-client.development.js:6863
	renderWithHooks @ react-dom-client.development.js:6775
	updateFunctionComponent @ react-dom-client.development.js:9069
	beginWork @ react-dom-client.development.js:10679
	runWithFiberInDEV @ react-dom-client.development.js:872
	performUnitOfWork @ react-dom-client.development.js:15677
	workLoopSync @ react-dom-client.development.js:15497
	renderRootSync @ react-dom-client.development.js:15477
	performWorkOnRoot @ react-dom-client.development.js:14941
	performWorkOnRootViaSchedulerTask @ react-dom-client.development.js:16766
	performWorkUntilDeadline @ scheduler.development.js:45
	<StoreBuilderPage>
	exports.jsx @ react-jsx-runtime.development.js:338
	ClientPageRoot @ client-page.tsx:60
	react_stack_bottom_frame @ react-dom-client.development.js:23552
	renderWithHooksAgain @ react-dom-client.development.js:6863
	renderWithHooks @ react-dom-client.development.js:6775
	updateFunctionComponent @ react-dom-client.development.js:9069
	beginWork @ react-dom-client.development.js:10628
	runWithFiberInDEV @ react-dom-client.development.js:872
	performUnitOfWork @ react-dom-client.development.js:15677
	workLoopConcurrentByScheduler @ react-dom-client.development.js:15671
	renderRootConcurrent @ react-dom-client.development.js:15646
	performWorkOnRoot @ react-dom-client.development.js:14940
	performWorkOnRootViaSchedulerTask @ react-dom-client.development.js:16766
	performWorkUntilDeadline @ scheduler.development.js:45
	"use client"
	initializeElement @ react-server-dom-turbopack-client.browser.development.js:1199
	wakeChunk @ react-server-dom-turbopack-client.browser.development.js:969
	fulfillReference @ react-server-dom-turbopack-client.browser.development.js:1303
	wakeChunk @ react-server-dom-turbopack-client.browser.development.js:970
	wakeChunkIfInitialized @ react-server-dom-turbopack-client.browser.development.js:1005
	resolveModuleChunk @ react-server-dom-turbopack-client.browser.development.js:1090
	(anonymous) @ react-server-dom-turbopack-client.browser.development.js:1937
	"use server"
	ResponseInstance @ react-server-dom-turbopack-client.browser.development.js:1863
	createResponseFromOptions @ react-server-dom-turbopack-client.browser.development.js:2851
	exports.createFromReadableStream @ react-server-dom-turbopack-client.browser.development.js:3213
	[project]/node_modules/next/dist/client/app-index.js [app-client] (ecmascript) @ app-index.tsx:157
	(anonymous) @ dev-base.ts:201
	runModuleExecutionHooks @ dev-base.ts:256
	instantiateModule @ dev-base.ts:199
	getOrInstantiateModuleFromParent @ dev-base.ts:126
	commonJsRequire @ runtime-utils.ts:291
	(anonymous) @ app-next-turbopack.ts:11
	(anonymous) @ app-bootstrap.ts:78
	loadScriptsInSequence @ app-bootstrap.ts:20
	appBootstrap @ app-bootstrap.ts:60
	[project]/node_modules/next/dist/client/app-next-turbopack.js [app-client] (ecmascript) @ app-next-turbopack.ts:10
	(anonymous) @ dev-base.ts:201
	runModuleExecutionHooks @ dev-base.ts:256
	instantiateModule @ dev-base.ts:199
	getOrInstantiateRuntimeModule @ dev-base.ts:96
	registerChunk @ runtime-backend-dom.ts:88
	await in registerChunk
	registerChunk @ runtime-base.ts:377
	(anonymous) @ dev-backend-dom.ts:126
	(anonymous) @ dev-backend-dom.ts:126
	component-library.tsx:496 Encountered two children with the same key, `product-showcase`. Keys should be unique so that components maintain their identity across updates. Non-unique keys may cause children to be duplicated and/or omitted — the behavior is unsupported and could change in a future version.
	overrideMethod @ hook.js:608
	error @ intercept-console-error.ts:44
	(anonymous) @ react-dom-client.development.js:5705
	runWithFiberInDEV @ react-dom-client.development.js:872
	warnOnInvalidKey @ react-dom-client.development.js:5704
	reconcileChildrenArray @ react-dom-client.development.js:5773
	reconcileChildFibersImpl @ react-dom-client.development.js:6094
	(anonymous) @ react-dom-client.development.js:6199
	reconcileChildren @ react-dom-client.development.js:8753
	beginWork @ react-dom-client.development.js:10950
	runWithFiberInDEV @ react-dom-client.development.js:872
	performUnitOfWork @ react-dom-client.development.js:15677
	workLoopSync @ react-dom-client.development.js:15497
	renderRootSync @ react-dom-client.development.js:15477
	performWorkOnRoot @ react-dom-client.development.js:14941
	performSyncWorkOnRoot @ react-dom-client.development.js:16781
	flushSyncWorkAcrossRoots_impl @ react-dom-client.development.js:16627
	processRootScheduleInMicrotask @ react-dom-client.development.js:16665
	(anonymous) @ react-dom-client.development.js:16800
	<div>
	exports.jsxDEV @ react-jsx-dev-runtime.development.js:345
	(anonymous) @ component-library.tsx:496
	ComponentLibrary @ component-library.tsx:495
	react_stack_bottom_frame @ react-dom-client.development.js:23552
	renderWithHooksAgain @ react-dom-client.development.js:6863
	renderWithHooks @ react-dom-client.development.js:6775
	updateFunctionComponent @ react-dom-client.development.js:9069
	beginWork @ react-dom-client.development.js:10679
	runWithFiberInDEV @ react-dom-client.development.js:872
	performUnitOfWork @ react-dom-client.development.js:15677
	workLoopSync @ react-dom-client.development.js:15497
	renderRootSync @ react-dom-client.development.js:15477
	performWorkOnRoot @ react-dom-client.development.js:14941
	performSyncWorkOnRoot @ react-dom-client.development.js:16781
	flushSyncWorkAcrossRoots_impl @ react-dom-client.development.js:16627
	processRootScheduleInMicrotask @ react-dom-client.development.js:16665
	(anonymous) @ react-dom-client.development.js:16800
	<ComponentLibrary>
	exports.jsxDEV @ react-jsx-dev-runtime.development.js:345
	StoreBuilder @ store-builder.tsx:2433
	react_stack_bottom_frame @ react-dom-client.development.js:23552
	renderWithHooksAgain @ react-dom-client.development.js:6863
	renderWithHooks @ react-dom-client.development.js:6775
	updateFunctionComponent @ react-dom-client.development.js:9069
	beginWork @ react-dom-client.development.js:10679
	runWithFiberInDEV @ react-dom-client.development.js:872
	performUnitOfWork @ react-dom-client.development.js:15677
	workLoopSync @ react-dom-client.development.js:15497
	renderRootSync @ react-dom-client.development.js:15477
	performWorkOnRoot @ react-dom-client.development.js:14941
	performSyncWorkOnRoot @ react-dom-client.development.js:16781
	flushSyncWorkAcrossRoots_impl @ react-dom-client.development.js:16627
	processRootScheduleInMicrotask @ react-dom-client.development.js:16665
	(anonymous) @ react-dom-client.development.js:16800
	<StoreBuilder>
	exports.jsxDEV @ react-jsx-dev-runtime.development.js:345
	StoreBuilderPage @ page.tsx:75
	react_stack_bottom_frame @ react-dom-client.development.js:23552
	renderWithHooksAgain @ react-dom-client.development.js:6863
	renderWithHooks @ react-dom-client.development.js:6775
	updateFunctionComponent @ react-dom-client.development.js:9069
	beginWork @ react-dom-client.development.js:10679
	runWithFiberInDEV @ react-dom-client.development.js:872
	performUnitOfWork @ react-dom-client.development.js:15677
	workLoopSync @ react-dom-client.development.js:15497
	renderRootSync @ react-dom-client.development.js:15477
	performWorkOnRoot @ react-dom-client.development.js:14941
	performWorkOnRootViaSchedulerTask @ react-dom-client.development.js:16766
	performWorkUntilDeadline @ scheduler.development.js:45
	<StoreBuilderPage>
	exports.jsx @ react-jsx-runtime.development.js:338
	ClientPageRoot @ client-page.tsx:60
	react_stack_bottom_frame @ react-dom-client.development.js:23552
	renderWithHooksAgain @ react-dom-client.development.js:6863
	renderWithHooks @ react-dom-client.development.js:6775
	updateFunctionComponent @ react-dom-client.development.js:9069
	beginWork @ react-dom-client.development.js:10628
	runWithFiberInDEV @ react-dom-client.development.js:872
	performUnitOfWork @ react-dom-client.development.js:15677
	workLoopConcurrentByScheduler @ react-dom-client.development.js:15671
	renderRootConcurrent @ react-dom-client.development.js:15646
	performWorkOnRoot @ react-dom-client.development.js:14940
	performWorkOnRootViaSchedulerTask @ react-dom-client.development.js:16766
	performWorkUntilDeadline @ scheduler.development.js:45
	"use client"
	initializeElement @ react-server-dom-turbopack-client.browser.development.js:1199
	wakeChunk @ react-server-dom-turbopack-client.browser.development.js:969
	fulfillReference @ react-server-dom-turbopack-client.browser.development.js:1303
	wakeChunk @ react-server-dom-turbopack-client.browser.development.js:970
	wakeChunkIfInitialized @ react-server-dom-turbopack-client.browser.development.js:1005
	resolveModuleChunk @ react-server-dom-turbopack-client.browser.development.js:1090
	(anonymous) @ react-server-dom-turbopack-client.browser.development.js:1937
	"use server"
	ResponseInstance @ react-server-dom-turbopack-client.browser.development.js:1863
	createResponseFromOptions @ react-server-dom-turbopack-client.browser.development.js:2851
	exports.createFromReadableStream @ react-server-dom-turbopack-client.browser.development.js:3213
	[project]/node_modules/next/dist/client/app-index.js [app-client] (ecmascript) @ app-index.tsx:157
	(anonymous) @ dev-base.ts:201
	runModuleExecutionHooks @ dev-base.ts:256
	instantiateModule @ dev-base.ts:199
	getOrInstantiateModuleFromParent @ dev-base.ts:126
	commonJsRequire @ runtime-utils.ts:291
	(anonymous) @ app-next-turbopack.ts:11
	(anonymous) @ app-bootstrap.ts:78
	loadScriptsInSequence @ app-bootstrap.ts:20
	appBootstrap @ app-bootstrap.ts:60
	[project]/node_modules/next/dist/client/app-next-turbopack.js [app-client] (ecmascript) @ app-next-turbopack.ts:10
	(anonymous) @ dev-base.ts:201
	runModuleExecutionHooks @ dev-base.ts:256
	instantiateModule @ dev-base.ts:199
	getOrInstantiateRuntimeModule @ dev-base.ts:96
	registerChunk @ runtime-backend-dom.ts:88
	await in registerChunk
	registerChunk @ runtime-base.ts:377
	(anonymous) @ dev-backend-dom.ts:126
	(anonymous) @ dev-backend-dom.ts:126

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

		// Newsletter subscriptions
		public.POST("/stores/:slug/newsletter/subscribe", newsletterController.Subscribe)
		public.GET("/stores/:slug/newsletter/unsubscribe", newsletterController.Unsubscribe)
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

			// Newsletter management
			storeRoutes.GET("/:id/newsletter/subscriptions", newsletterController.GetSubscriptions)
			storeRoutes.DELETE("/:id/newsletter/subscriptions/:subscriptionId", newsletterController.DeleteSubscription)
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

package router

import (
	"admin-backend/internal/handler"
	"admin-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORS())

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api := r.Group("/api/v1")

	// Auth routes (no auth required for login)
	authHandler := handler.NewAuthHandler()
	auth := api.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
	}

	// Protected routes
	protected := api.Group("")
	protected.Use(middleware.AdminAuth())
	protected.Use(middleware.OperationLogger())

	// Auth routes (auth required)
	protectedAuth := protected.Group("/auth")
	{
		protectedAuth.POST("/logout", authHandler.Logout)
		protectedAuth.PUT("/password", authHandler.ChangePassword)
		protectedAuth.GET("/profile", authHandler.GetProfile)
	}

	// Dashboard
	dashboardHandler := handler.NewDashboardHandler()
	dashboard := protected.Group("/dashboard")
	{
		dashboard.GET("/stats", dashboardHandler.GetStats)
		dashboard.GET("/download-trend", dashboardHandler.GetDownloadTrend)
		dashboard.GET("/user-trend", dashboardHandler.GetUserTrend)
		dashboard.GET("/category-distribution", dashboardHandler.GetCategoryDistribution)
		dashboard.GET("/top-software", dashboardHandler.GetTopSoftware)
		dashboard.GET("/todos", dashboardHandler.GetTodos)
	}

	// Software management
	softwareHandler := handler.NewSoftwareHandler()
	software := protected.Group("/software")
	{
		software.GET("", softwareHandler.List)
		software.GET("/:id", softwareHandler.GetByID)
		software.PUT("/:id", middleware.RequirePermission("software:edit"), softwareHandler.Update)
		software.PUT("/:id/on-shelf", middleware.RequirePermission("software:shelf"), softwareHandler.OnShelf)
		software.PUT("/:id/off-shelf", middleware.RequirePermission("software:shelf"), softwareHandler.OffShelf)
		software.DELETE("/:id", middleware.RequirePermission("software:delete"), softwareHandler.Delete)
		software.POST("/batch", middleware.RequirePermission("software:batch"), softwareHandler.BatchOperation)
	}

	// Category management
	categoryHandler := handler.NewCategoryHandler()
	categories := protected.Group("/categories")
	{
		categories.GET("", categoryHandler.List)
		categories.POST("", middleware.RequirePermission("category:create"), categoryHandler.Create)
		categories.PUT("/:id", middleware.RequirePermission("category:edit"), categoryHandler.Update)
		categories.PUT("/:id/status", middleware.RequirePermission("category:edit"), categoryHandler.ToggleStatus)
		categories.DELETE("/:id", middleware.RequirePermission("category:delete"), categoryHandler.Delete)
	}

	// Tag management
	tagHandler := handler.NewTagHandler()
	tags := protected.Group("/tags")
	{
		tags.GET("", tagHandler.List)
		tags.POST("", middleware.RequirePermission("tag:create"), tagHandler.Create)
		tags.PUT("/:id", middleware.RequirePermission("tag:edit"), tagHandler.Update)
		tags.PUT("/:id/status", middleware.RequirePermission("tag:edit"), tagHandler.ToggleStatus)
		tags.DELETE("/:id", middleware.RequirePermission("tag:delete"), tagHandler.Delete)
	}

	// Banner management
	bannerHandler := handler.NewBannerHandler()
	banners := protected.Group("/banners")
	{
		banners.GET("", bannerHandler.List)
		banners.POST("", middleware.RequirePermission("banner:create"), bannerHandler.Create)
		banners.PUT("/sort", middleware.RequirePermission("banner:edit"), bannerHandler.BatchSort)
		banners.PUT("/:id", middleware.RequirePermission("banner:edit"), bannerHandler.Update)
		banners.PUT("/:id/status", middleware.RequirePermission("banner:edit"), bannerHandler.ToggleStatus)
		banners.DELETE("/:id", middleware.RequirePermission("banner:delete"), bannerHandler.Delete)
	}

	// Recommendation management
	recommendationHandler := handler.NewRecommendationHandler()
	recommendations := protected.Group("/recommendations")
	{
		recommendations.GET("", recommendationHandler.List)
		recommendations.POST("", middleware.RequirePermission("recommendation:create"), recommendationHandler.Create)
		recommendations.PUT("/sort", middleware.RequirePermission("recommendation:edit"), recommendationHandler.BatchSort)
		recommendations.PUT("/:id", middleware.RequirePermission("recommendation:edit"), recommendationHandler.Update)
		recommendations.DELETE("/:id", middleware.RequirePermission("recommendation:delete"), recommendationHandler.Delete)
	}

	// Announcement management
	announcementHandler := handler.NewAnnouncementHandler()
	announcements := protected.Group("/announcements")
	{
		announcements.GET("", announcementHandler.List)
		announcements.POST("", middleware.RequirePermission("announcement:create"), announcementHandler.Create)
		announcements.PUT("/:id", middleware.RequirePermission("announcement:edit"), announcementHandler.Update)
		announcements.PUT("/:id/publish", middleware.RequirePermission("announcement:publish"), announcementHandler.Publish)
		announcements.PUT("/:id/withdraw", middleware.RequirePermission("announcement:publish"), announcementHandler.Withdraw)
		announcements.DELETE("/:id", middleware.RequirePermission("announcement:delete"), announcementHandler.Delete)
	}

	// Audit management
	auditHandler := handler.NewAuditHandler()
	audits := protected.Group("/audits")
	{
		audits.GET("", auditHandler.List)
		audits.GET("/:id", auditHandler.GetByID)
		audits.POST("/:id/approve", middleware.RequirePermission("audit:review"), auditHandler.Approve)
		audits.POST("/:id/reject", middleware.RequirePermission("audit:review"), auditHandler.Reject)
	}

	// Review audit
	reviewAuditHandler := handler.NewReviewAuditHandler()
	reviewAudits := protected.Group("/review-audits")
	{
		reviewAudits.GET("", reviewAuditHandler.List)
		reviewAudits.PUT("/:id/hide", middleware.RequirePermission("review:moderate"), reviewAuditHandler.Hide)
		reviewAudits.PUT("/:id/restore", middleware.RequirePermission("review:moderate"), reviewAuditHandler.Restore)
		reviewAudits.DELETE("/:id", middleware.RequirePermission("review:delete"), reviewAuditHandler.Delete)
	}

	// Report management
	reportHandler := handler.NewReportHandler()
	reports := protected.Group("/reports")
	{
		reports.GET("", reportHandler.List)
		reports.GET("/:id", reportHandler.GetByID)
		reports.PUT("/:id/resolve", middleware.RequirePermission("report:resolve"), reportHandler.Resolve)
	}

	// User management
	userHandler := handler.NewUserHandler()
	users := protected.Group("/users")
	{
		users.GET("", userHandler.List)
		users.GET("/:id", userHandler.GetDetail)
		users.PUT("/:id/disable", middleware.RequirePermission("user:disable"), userHandler.Disable)
		users.PUT("/:id/enable", middleware.RequirePermission("user:enable"), userHandler.Enable)
		users.PUT("/:id/space", middleware.RequirePermission("user:space"), userHandler.AdjustSpace)
	}

	// Feedback management
	feedbackHandler := handler.NewFeedbackHandler()
	feedbacks := protected.Group("/feedbacks")
	{
		feedbacks.GET("", feedbackHandler.List)
		feedbacks.GET("/:id", feedbackHandler.GetByID)
		feedbacks.PUT("/:id/assign", middleware.RequirePermission("feedback:assign"), feedbackHandler.Assign)
		feedbacks.PUT("/:id/reply", middleware.RequirePermission("feedback:reply"), feedbackHandler.Reply)
		feedbacks.PUT("/:id/close", middleware.RequirePermission("feedback:close"), feedbackHandler.Close)
	}

	// Stats
	statsHandler := handler.NewStatsHandler()
	stats := protected.Group("/stats")
	{
		stats.GET("/downloads", statsHandler.GetDownloadStats)
		stats.GET("/users", statsHandler.GetUserStats)
		stats.GET("/software-ranking", statsHandler.GetSoftwareRanking)
	}

	// Admin management
	adminHandler := handler.NewAdminHandler()
	admins := protected.Group("/admins")
	{
		admins.GET("", middleware.RequirePermission("admin:view"), adminHandler.List)
		admins.POST("", middleware.RequirePermission("admin:create"), adminHandler.Create)
		admins.PUT("/:id", middleware.RequirePermission("admin:edit"), adminHandler.Update)
		admins.PUT("/:id/reset-password", middleware.RequirePermission("admin:reset_password"), adminHandler.ResetPassword)
		admins.PUT("/:id/status", middleware.RequirePermission("admin:edit"), adminHandler.ToggleStatus)
		admins.DELETE("/:id", middleware.RequirePermission("admin:delete"), adminHandler.Delete)
	}

	// Role & Permission management
	roleHandler := handler.NewRoleHandler()
	roles := protected.Group("/roles")
	{
		roles.GET("", roleHandler.List)
		roles.POST("", middleware.RequirePermission("role:create"), roleHandler.Create)
		roles.PUT("/:id", middleware.RequirePermission("role:edit"), roleHandler.Update)
		roles.DELETE("/:id", middleware.RequirePermission("role:delete"), roleHandler.Delete)
		roles.PUT("/:id/permissions", middleware.RequirePermission("role:set_permissions"), roleHandler.SetPermissions)
	}
	permissions := protected.Group("/permissions")
	{
		permissions.GET("", roleHandler.ListPermissions)
	}

	// Operation logs
	operationLogHandler := handler.NewOperationLogHandler()
	operationLogs := protected.Group("/operation-logs")
	{
		operationLogs.GET("", operationLogHandler.List)
	}

	// System configs
	systemConfigHandler := handler.NewSystemConfigHandler()
	systemConfigs := protected.Group("/system-configs")
	{
		systemConfigs.GET("", systemConfigHandler.ListAll)
		systemConfigs.PUT("", middleware.RequirePermission("system:config"), systemConfigHandler.BatchUpdate)
	}

	return r
}

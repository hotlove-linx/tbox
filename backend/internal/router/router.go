package router

import (
	"tbox-backend/internal/handler"
	"tbox-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(mode string) *gin.Engine {
	gin.SetMode(mode)
	r := gin.New()

	// Global middleware
	r.Use(middleware.LoggerMiddleware())
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())

	// Serve uploaded files
	r.Static("/uploads", "./uploads")

	// API v1
	v1 := r.Group("/api/v1")

	// Handlers
	authHandler := handler.NewAuthHandler()
	userHandler := handler.NewUserHandler()
	softwareHandler := handler.NewSoftwareHandler()
	reviewHandler := handler.NewReviewHandler()
	spaceHandler := handler.NewSpaceHandler()

	// Auth routes (public)
	auth := v1.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/send-code", authHandler.SendCode)
		auth.POST("/reset-password", authHandler.ResetPassword)
	}

	// Auth logout (requires auth)
	authProtected := v1.Group("/auth")
	authProtected.Use(middleware.AuthMiddleware())
	{
		authProtected.POST("/logout", authHandler.Logout)
	}

	// User routes (requires auth)
	user := v1.Group("/user")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/profile", userHandler.GetProfile)
		user.PUT("/profile", userHandler.UpdateProfile)
		user.PUT("/avatar", userHandler.UpdateAvatar)
		user.PUT("/password", userHandler.ChangePassword)
		user.DELETE("/account", userHandler.DeleteAccount)
	}

	// Software routes (public)
	software := v1.Group("/software")
	{
		software.GET("/home", softwareHandler.GetHome)
		software.GET("/list", softwareHandler.GetList)
		software.GET("/search", softwareHandler.Search)
		software.GET("/:id", softwareHandler.GetDetail)
		software.GET("/:id/versions", softwareHandler.GetVersions)
		software.GET("/:id/reviews", reviewHandler.GetReviews)
		software.GET("/:id/download", softwareHandler.GetDownloadURL)
		software.POST("/:id/download-count", softwareHandler.RecordDownload)
		software.GET("/:id/check-update", softwareHandler.CheckUpdate)
	}

	// Categories (public)
	v1.GET("/categories", softwareHandler.GetCategories)

	// Reviews (requires auth)
	reviews := v1.Group("")
	reviews.Use(middleware.AuthMiddleware())
	{
		reviews.POST("/software/:id/reviews", reviewHandler.CreateReview)
		reviews.PUT("/reviews/:id", reviewHandler.UpdateReview)
		reviews.DELETE("/reviews/:id", reviewHandler.DeleteReview)
	}

	// Space routes (requires auth)
	space := v1.Group("/space")
	space.Use(middleware.AuthMiddleware())
	{
		space.GET("/overview", spaceHandler.GetOverview)
		space.GET("/software", spaceHandler.GetMySoftware)
		space.POST("/software", spaceHandler.CreateSoftware)
		space.PUT("/software/:id", spaceHandler.UpdateSoftware)
		space.POST("/software/:id/submit-audit", spaceHandler.SubmitAudit)
		space.POST("/software/:id/withdraw", spaceHandler.WithdrawAudit)
	}

	// Upload routes (requires auth)
	upload := v1.Group("/upload")
	upload.Use(middleware.AuthMiddleware())
	{
		upload.POST("/init", spaceHandler.InitUpload)
		upload.POST("/chunk", spaceHandler.UploadChunk)
		upload.POST("/complete", spaceHandler.CompleteUpload)
	}

	return r
}

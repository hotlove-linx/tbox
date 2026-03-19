package main

import (
	"fmt"
	"log"

	"admin-backend/config"
	"admin-backend/internal/model"
	"admin-backend/internal/router"
	"admin-backend/pkg/database"
	"admin-backend/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Set gin mode
	gin.SetMode(cfg.Server.Mode)

	// Init MySQL
	if err := database.InitMySQL(&cfg.Database); err != nil {
		log.Fatalf("Failed to init MySQL: %v", err)
	}

	// Init Redis
	if err := database.InitRedis(&cfg.Redis); err != nil {
		log.Fatalf("Failed to init Redis: %v", err)
	}

	// Auto migrate admin-specific tables (shared tables are managed by backend)
	autoMigrate(database.GetDB())

	// Seed default data
	seedData(database.GetDB(), cfg)

	// Setup router
	r := router.SetupRouter()

	// Start server
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Admin backend starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Admin{},
		&model.Role{},
		&model.AdminRole{},
		&model.Permission{},
		&model.RolePermission{},
		&model.OperationLog{},
		&model.Banner{},
		&model.Announcement{},
		&model.Recommendation{},
		&model.Report{},
		&model.Feedback{},
		&model.SystemConfig{},
	)
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
	log.Println("Database migration completed")
}

func seedData(db *gorm.DB, cfg *config.Config) {
	// Seed permissions
	seedPermissions(db)

	// Seed roles
	seedRoles(db)

	// Seed default admin
	seedDefaultAdmin(db, cfg)

	// Seed default system configs
	seedSystemConfigs(db)

	log.Println("Seed data initialized")
}

func seedPermissions(db *gorm.DB) {
	permissions := []model.Permission{
		// Software
		{Module: "software", Action: "view", Code: "software:view", Description: "View software"},
		{Module: "software", Action: "edit", Code: "software:edit", Description: "Edit software"},
		{Module: "software", Action: "delete", Code: "software:delete", Description: "Delete software"},
		{Module: "software", Action: "shelf", Code: "software:shelf", Description: "On/Off shelf software"},
		{Module: "software", Action: "batch", Code: "software:batch", Description: "Batch operations on software"},
		// Category
		{Module: "category", Action: "create", Code: "category:create", Description: "Create category"},
		{Module: "category", Action: "edit", Code: "category:edit", Description: "Edit category"},
		{Module: "category", Action: "delete", Code: "category:delete", Description: "Delete category"},
		// Tag
		{Module: "tag", Action: "create", Code: "tag:create", Description: "Create tag"},
		{Module: "tag", Action: "edit", Code: "tag:edit", Description: "Edit tag"},
		{Module: "tag", Action: "delete", Code: "tag:delete", Description: "Delete tag"},
		// Banner
		{Module: "banner", Action: "create", Code: "banner:create", Description: "Create banner"},
		{Module: "banner", Action: "edit", Code: "banner:edit", Description: "Edit banner"},
		{Module: "banner", Action: "delete", Code: "banner:delete", Description: "Delete banner"},
		// Recommendation
		{Module: "recommendation", Action: "create", Code: "recommendation:create", Description: "Create recommendation"},
		{Module: "recommendation", Action: "edit", Code: "recommendation:edit", Description: "Edit recommendation"},
		{Module: "recommendation", Action: "delete", Code: "recommendation:delete", Description: "Delete recommendation"},
		// Announcement
		{Module: "announcement", Action: "create", Code: "announcement:create", Description: "Create announcement"},
		{Module: "announcement", Action: "edit", Code: "announcement:edit", Description: "Edit announcement"},
		{Module: "announcement", Action: "publish", Code: "announcement:publish", Description: "Publish/withdraw announcement"},
		{Module: "announcement", Action: "delete", Code: "announcement:delete", Description: "Delete announcement"},
		// Audit
		{Module: "audit", Action: "review", Code: "audit:review", Description: "Review software audit"},
		// Review
		{Module: "review", Action: "moderate", Code: "review:moderate", Description: "Moderate reviews"},
		{Module: "review", Action: "delete", Code: "review:delete", Description: "Delete reviews"},
		// Report
		{Module: "report", Action: "resolve", Code: "report:resolve", Description: "Resolve reports"},
		// User
		{Module: "user", Action: "disable", Code: "user:disable", Description: "Disable user"},
		{Module: "user", Action: "enable", Code: "user:enable", Description: "Enable user"},
		{Module: "user", Action: "space", Code: "user:space", Description: "Adjust user space"},
		// Feedback
		{Module: "feedback", Action: "assign", Code: "feedback:assign", Description: "Assign feedback"},
		{Module: "feedback", Action: "reply", Code: "feedback:reply", Description: "Reply feedback"},
		{Module: "feedback", Action: "close", Code: "feedback:close", Description: "Close feedback"},
		// Admin
		{Module: "admin", Action: "view", Code: "admin:view", Description: "View admins"},
		{Module: "admin", Action: "create", Code: "admin:create", Description: "Create admin"},
		{Module: "admin", Action: "edit", Code: "admin:edit", Description: "Edit admin"},
		{Module: "admin", Action: "delete", Code: "admin:delete", Description: "Delete admin"},
		{Module: "admin", Action: "reset_password", Code: "admin:reset_password", Description: "Reset admin password"},
		// Role
		{Module: "role", Action: "create", Code: "role:create", Description: "Create role"},
		{Module: "role", Action: "edit", Code: "role:edit", Description: "Edit role"},
		{Module: "role", Action: "delete", Code: "role:delete", Description: "Delete role"},
		{Module: "role", Action: "set_permissions", Code: "role:set_permissions", Description: "Set role permissions"},
		// System
		{Module: "system", Action: "config", Code: "system:config", Description: "Manage system configs"},
	}

	for _, p := range permissions {
		var existing model.Permission
		if db.Where("code = ?", p.Code).First(&existing).Error != nil {
			db.Create(&p)
		}
	}
}

func seedRoles(db *gorm.DB) {
	roles := []model.Role{
		{Name: "Super Admin", Code: "super_admin", Description: "Full system access", IsSystem: true},
		{Name: "Content Manager", Code: "content_manager", Description: "Manage content and software", IsSystem: true},
		{Name: "Auditor", Code: "auditor", Description: "Review and audit software", IsSystem: true},
		{Name: "Customer Service", Code: "customer_service", Description: "Handle user feedback and reports", IsSystem: true},
	}

	for _, r := range roles {
		var existing model.Role
		if db.Where("code = ?", r.Code).First(&existing).Error != nil {
			db.Create(&r)
		}
	}

	// Assign all permissions to super_admin role
	var superAdminRole model.Role
	if db.Where("code = ?", "super_admin").First(&superAdminRole).Error == nil {
		var allPermissions []model.Permission
		db.Find(&allPermissions)

		for _, p := range allPermissions {
			var existing model.RolePermission
			if db.Where("role_id = ? AND permission_id = ?", superAdminRole.ID, p.ID).First(&existing).Error != nil {
				db.Create(&model.RolePermission{RoleID: superAdminRole.ID, PermissionID: p.ID})
			}
		}
	}
}

func seedDefaultAdmin(db *gorm.DB, cfg *config.Config) {
	var admin model.Admin
	if db.Where("email = ?", cfg.Admin.DefaultEmail).First(&admin).Error != nil {
		hash, err := utils.HashPassword(cfg.Admin.DefaultPassword)
		if err != nil {
			log.Printf("Failed to hash default admin password: %v", err)
			return
		}

		admin = model.Admin{
			Name:         "Super Admin",
			Email:        cfg.Admin.DefaultEmail,
			PasswordHash: hash,
			Status:       "active",
		}
		if err := db.Create(&admin).Error; err != nil {
			log.Printf("Failed to create default admin: %v", err)
			return
		}

		// Assign super_admin role
		var superAdminRole model.Role
		if db.Where("code = ?", "super_admin").First(&superAdminRole).Error == nil {
			db.Create(&model.AdminRole{AdminID: admin.ID, RoleID: superAdminRole.ID})
		}

		log.Printf("Default admin created: %s / %s", cfg.Admin.DefaultEmail, cfg.Admin.DefaultPassword)
	}
}

func seedSystemConfigs(db *gorm.DB) {
	configs := []model.SystemConfig{
		{ConfigGroup: "site", ConfigKey: "site_name", ConfigValue: "TBox Software Store", Description: "Site name"},
		{ConfigGroup: "site", ConfigKey: "site_description", ConfigValue: "Your trusted software marketplace", Description: "Site description"},
		{ConfigGroup: "upload", ConfigKey: "max_file_size", ConfigValue: "1073741824", Description: "Max upload file size in bytes (1GB)"},
		{ConfigGroup: "upload", ConfigKey: "allowed_extensions", ConfigValue: ".exe,.msi,.zip,.dmg,.pkg,.deb,.rpm,.appimage", Description: "Allowed file extensions"},
		{ConfigGroup: "user", ConfigKey: "default_space", ConfigValue: "5368709120", Description: "Default user space in bytes (5GB)"},
		{ConfigGroup: "audit", ConfigKey: "auto_approve", ConfigValue: "false", Description: "Auto approve software"},
	}

	for _, c := range configs {
		var existing model.SystemConfig
		if db.Where("config_key = ?", c.ConfigKey).First(&existing).Error != nil {
			db.Create(&c)
		}
	}
}

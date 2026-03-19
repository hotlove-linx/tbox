package middleware

import (
	"admin-backend/internal/model"
	"admin-backend/pkg/database"
	"admin-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

func RequirePermission(permissionCode string) gin.HandlerFunc {
	return func(c *gin.Context) {
		adminID := GetAdminID(c)
		if adminID == 0 {
			response.Unauthorized(c, "not authenticated")
			c.Abort()
			return
		}

		// Check if admin has the required permission through their roles
		db := database.GetDB()

		var count int64
		db.Table("role_permissions").
			Joins("JOIN admin_roles ON admin_roles.role_id = role_permissions.role_id").
			Joins("JOIN permissions ON permissions.id = role_permissions.permission_id").
			Where("admin_roles.admin_id = ? AND permissions.code = ?", adminID, permissionCode).
			Count(&count)

		if count == 0 {
			// Check if admin has super_admin role (bypass permission check)
			var superCount int64
			db.Model(&model.AdminRole{}).
				Joins("JOIN roles ON roles.id = admin_roles.role_id").
				Where("admin_roles.admin_id = ? AND roles.code = ?", adminID, "super_admin").
				Count(&superCount)

			if superCount == 0 {
				response.Forbidden(c, "insufficient permissions")
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

package middleware

import (
	"context"
	"fmt"
	"strings"

	"admin-backend/pkg/database"
	"admin-backend/pkg/response"
	"admin-backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

const (
	AdminSessionPrefix = "admin:session:"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Unauthorized(c, "missing authorization header")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Unauthorized(c, "invalid authorization format")
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := utils.ParseAdminToken(tokenString)
		if err != nil {
			response.Unauthorized(c, "invalid or expired token")
			c.Abort()
			return
		}

		// Check session in Redis
		rdb := database.GetRedis()
		sessionKey := fmt.Sprintf("%s%d", AdminSessionPrefix, claims.AdminID)
		storedToken, err := rdb.Get(context.Background(), sessionKey).Result()
		if err != nil || storedToken != tokenString {
			response.Unauthorized(c, "session expired or invalid")
			c.Abort()
			return
		}

		c.Set("admin_id", claims.AdminID)
		c.Set("admin_email", claims.Email)
		c.Next()
	}
}

func GetAdminID(c *gin.Context) uint64 {
	adminID, exists := c.Get("admin_id")
	if !exists {
		return 0
	}
	return adminID.(uint64)
}

func GetAdminEmail(c *gin.Context) string {
	email, exists := c.Get("admin_email")
	if !exists {
		return ""
	}
	return email.(string)
}

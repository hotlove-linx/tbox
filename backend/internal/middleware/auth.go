package middleware

import (
	"context"
	"strings"

	"tbox-backend/pkg/database"
	"tbox-backend/pkg/response"
	"tbox-backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Unauthorized(c, "请先登录")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Unauthorized(c, "认证格式错误")
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			response.Unauthorized(c, "认证已过期，请重新登录")
			c.Abort()
			return
		}

		// Check if token exists in Redis (session validation)
		ctx := context.Background()
		rdb := database.GetRedis()
		key := "user:session:" + tokenString
		exists, err := rdb.Exists(ctx, key).Result()
		if err != nil || exists == 0 {
			response.Unauthorized(c, "登录已失效，请重新登录")
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("token", tokenString)
		c.Next()
	}
}

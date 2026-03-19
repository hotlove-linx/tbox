package utils

import (
	"github.com/gin-gonic/gin"
)

func GetUserID(c *gin.Context) uint64 {
	val, exists := c.Get("user_id")
	if !exists {
		return 0
	}
	if id, ok := val.(uint64); ok {
		return id
	}
	return 0
}

func GetToken(c *gin.Context) string {
	val, exists := c.Get("token")
	if !exists {
		return ""
	}
	if token, ok := val.(string); ok {
		return token
	}
	return ""
}

package middleware

import (
	"encoding/json"
	"log"
	"strings"

	"admin-backend/internal/model"
	"admin-backend/pkg/database"

	"github.com/gin-gonic/gin"
)

// moduleActionMap maps URL path patterns to module and action for operation logging
var moduleActionMap = map[string]struct {
	Module string
	Action string
}{
	"POST /api/v1/software/batch":      {Module: "software", Action: "batch"},
	"PUT /api/v1/software/":            {Module: "software", Action: "update"},
	"DELETE /api/v1/software/":         {Module: "software", Action: "delete"},
	"POST /api/v1/categories":          {Module: "category", Action: "create"},
	"PUT /api/v1/categories/":          {Module: "category", Action: "update"},
	"DELETE /api/v1/categories/":       {Module: "category", Action: "delete"},
	"POST /api/v1/tags":                {Module: "tag", Action: "create"},
	"PUT /api/v1/tags/":                {Module: "tag", Action: "update"},
	"DELETE /api/v1/tags/":             {Module: "tag", Action: "delete"},
	"POST /api/v1/banners":             {Module: "banner", Action: "create"},
	"PUT /api/v1/banners/":             {Module: "banner", Action: "update"},
	"DELETE /api/v1/banners/":          {Module: "banner", Action: "delete"},
	"POST /api/v1/recommendations":     {Module: "recommendation", Action: "create"},
	"PUT /api/v1/recommendations/":     {Module: "recommendation", Action: "update"},
	"DELETE /api/v1/recommendations/":  {Module: "recommendation", Action: "delete"},
	"POST /api/v1/announcements":       {Module: "announcement", Action: "create"},
	"PUT /api/v1/announcements/":       {Module: "announcement", Action: "update"},
	"DELETE /api/v1/announcements/":    {Module: "announcement", Action: "delete"},
	"POST /api/v1/audits/":             {Module: "audit", Action: "review"},
	"PUT /api/v1/review-audits/":       {Module: "review_audit", Action: "update"},
	"DELETE /api/v1/review-audits/":    {Module: "review_audit", Action: "delete"},
	"PUT /api/v1/reports/":             {Module: "report", Action: "resolve"},
	"PUT /api/v1/users/":               {Module: "user", Action: "update"},
	"PUT /api/v1/feedbacks/":           {Module: "feedback", Action: "update"},
	"POST /api/v1/admins":              {Module: "admin", Action: "create"},
	"PUT /api/v1/admins/":              {Module: "admin", Action: "update"},
	"DELETE /api/v1/admins/":           {Module: "admin", Action: "delete"},
	"POST /api/v1/roles":               {Module: "role", Action: "create"},
	"PUT /api/v1/roles/":               {Module: "role", Action: "update"},
	"DELETE /api/v1/roles/":            {Module: "role", Action: "delete"},
	"PUT /api/v1/system-configs":       {Module: "system_config", Action: "update"},
}

func OperationLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Only log write operations
		method := c.Request.Method
		if method == "GET" {
			c.Next()
			return
		}

		c.Next()

		// Only log successful operations
		if c.Writer.Status() >= 400 {
			return
		}

		adminID := GetAdminID(c)
		if adminID == 0 {
			return
		}

		path := c.Request.URL.Path
		var module, action string

		// Try exact match first, then prefix match
		key := method + " " + path
		if info, ok := moduleActionMap[key]; ok {
			module = info.Module
			action = info.Action
		} else {
			for pattern, info := range moduleActionMap {
				if strings.HasPrefix(key, pattern) {
					module = info.Module
					action = info.Action
					break
				}
			}
		}

		if module == "" {
			return
		}

		// Enhance action based on path suffix
		if strings.HasSuffix(path, "/on-shelf") {
			action = "on_shelf"
		} else if strings.HasSuffix(path, "/off-shelf") {
			action = "off_shelf"
		} else if strings.HasSuffix(path, "/status") {
			action = "toggle_status"
		} else if strings.HasSuffix(path, "/approve") {
			action = "approve"
		} else if strings.HasSuffix(path, "/reject") {
			action = "reject"
		} else if strings.HasSuffix(path, "/publish") {
			action = "publish"
		} else if strings.HasSuffix(path, "/withdraw") {
			action = "withdraw"
		} else if strings.HasSuffix(path, "/enable") {
			action = "enable"
		} else if strings.HasSuffix(path, "/disable") {
			action = "disable"
		} else if strings.HasSuffix(path, "/reset-password") {
			action = "reset_password"
		} else if strings.HasSuffix(path, "/permissions") {
			action = "set_permissions"
		}

		detail := map[string]interface{}{
			"method": method,
			"path":   path,
		}
		detailJSON, _ := json.Marshal(detail)

		opLog := model.OperationLog{
			AdminID: adminID,
			Module:  module,
			Action:  action,
			Detail:  string(detailJSON),
			IP:      c.ClientIP(),
		}

		if err := database.GetDB().Create(&opLog).Error; err != nil {
			log.Printf("Failed to create operation log: %v", err)
		}
	}
}

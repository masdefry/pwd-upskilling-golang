package middlewares

import (
	"intro-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizeRole(allowedRoles ...models.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Role not found",
			})
			return
		}

		for _, r := range allowedRoles {
			if role == r {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "Unathorized role",
		})
	}
}

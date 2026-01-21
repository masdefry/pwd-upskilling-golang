package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RecoverMiddleware menangkap panic dan mengembalikan response JSON
func RecoverMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"error":   "Internal Server Error",
					"message": r,
				})
				c.Abort() // Menghentikan request lebih lanjut
			}
		}()
		c.Next() 
	}
}

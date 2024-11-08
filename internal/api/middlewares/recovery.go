package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error":   "there are some internal error",
					"message": err,
				})
			}
		}()
		c.Next()
	}
}

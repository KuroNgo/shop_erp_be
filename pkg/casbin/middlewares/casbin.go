package middlewares

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Authorize determines if current user has been authorized to take an action on an object.
func Authorize(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser, exists := c.Get("currentUser")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": "You are not logged in!",
			})
			return
		}

		// Load policy from Database
		err := enforcer.LoadPolicy()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to load policy from DB",
			})
			return
		}

		// Casbin enforces policy
		object := "http://localhost:8080" + c.Request.URL.Path
		action := c.Request.Method
		ok, err := enforcer.Enforce(fmt.Sprint(currentUser), object, action)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Error occurred when authorizing user",
			})
			return
		}

		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "You are not authorized!",
			})
			return
		}
		c.Next()
	}
}

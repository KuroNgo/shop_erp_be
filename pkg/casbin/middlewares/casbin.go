package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_erp_mono/infrastructor"
	"shop_erp_mono/pkg/token"
)

// Authorize determines if current user has been authorized to take an action on an object.
func Authorize(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("access_token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "You are not login!",
			})
			return
		}

		app, _ := infrastructor.App()
		env := app.Env

		sub, err := token.ValidateToken(cookie, env.AccessTokenPublicKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
			return
		}

		// Load policy from Database
		err = enforcer.LoadPolicy()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to load policy from DB",
			})
			return
		}

		// Casbin enforces policy
		object := "http://localhost:8080" + c.Request.URL.Path
		action := c.Request.Method
		ok, err := enforcer.Enforce(fmt.Sprint(sub), object, action)

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

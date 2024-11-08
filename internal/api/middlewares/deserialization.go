package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_erp_mono/internal/infrastructor"
	"shop_erp_mono/pkg/shared/token"
	"strings"
)

func DeserializeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var accessToken string
		cookie, err := ctx.Cookie("access_token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			accessToken = fields[1]
		} else if err == nil {
			accessToken = cookie
		}

		if accessToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": "You are not logged in",
			})
			return
		}

		app, _ := infrastructor.App()
		env := app.Env

		sub, err := token.ValidateToken(accessToken, env.AccessTokenPublicKey)
		if err != nil {
			fmt.Println("The err is: ", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
			return
		}

		ctx.Set("currentUser", sub)
		ctx.Next() // Cho phép tiếp tục các handler khác nếu không có lỗi
	}
}

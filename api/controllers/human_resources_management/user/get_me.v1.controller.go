package user_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_erp_mono/pkg/token"
)

// GetMe retrieves the user information based on the access token.
// @Summary Get User Information
// @Description Retrieves the user's information using the access token stored in cookies.
// @Tags User
// @Accept  json
// @Produce  json
// @Router /api/v1/users/get/info [get]
// @Security CookieAuth
func (u *UserController) GetMe(ctx *gin.Context) {
	cookie, err := ctx.Cookie("access_token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "fail",
			"message": "You are not logged in!",
		})
		return
	}

	sub, err := token.ValidateToken(cookie, u.Database.AccessTokenPublicKey)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status":  "fail",
			"message": "Failed to validate token: " + err.Error(),
		})
		return
	}

	result, err := u.UserUseCase.GetByID(ctx, fmt.Sprint(sub))
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status":  "fail",
			"message": "Failed to get user data: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"user":   result,
	})
}

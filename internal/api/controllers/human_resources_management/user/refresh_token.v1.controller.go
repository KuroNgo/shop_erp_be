package user_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RefreshToken refreshes the user's access token.
// @Summary Refresh Access Token
// @Description Refresh the user's access token using a valid refresh token stored in cookies.
// @Tags User
// @Accept  json
// @Produce  json
// @Router /api/v1/users/get/refresh [get]
// @Security CookieAuth
func (u *UserController) RefreshToken(ctx *gin.Context) {
	cookie, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status":  "fail",
			"message": "could not refresh access token",
		})
		return
	}

	data, err := u.UserUseCase.RefreshToken(ctx, cookie)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": "could not refresh access token",
		})
		return
	}

	ctx.SetCookie("access_token", data.AccessToken, u.Database.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", data.RefreshToken, u.Database.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", u.Database.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{
		"status":        "success",
		"access_token":  data.AccessToken,
		"refresh_token": data.RefreshToken,
	})
}

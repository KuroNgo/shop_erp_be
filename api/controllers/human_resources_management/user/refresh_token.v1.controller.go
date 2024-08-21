package user_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_erp_mono/pkg/token"
)

func (u *UserController) RefreshToken(ctx *gin.Context) {
	message := "could not refresh access token"

	cookie, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status":  "fail",
			"message": message,
		})
		return
	}

	sub, err := token.ValidateToken(cookie, u.Database.RefreshTokenPublicKey)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	user, err := u.UserUseCase.GetByID(ctx, fmt.Sprint(sub))
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status":  "fail",
			"message": "the user belonging to this token no logger exists",
		})
		return
	}

	access_token, err := token.CreateToken(u.Database.AccessTokenExpiresIn, user.ID, u.Database.AccessTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	ctx.SetCookie("access_token", access_token, u.Database.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", u.Database.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{
		"status":       "success",
		"access_token": access_token,
	})
}

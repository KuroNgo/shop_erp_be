package user_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GoogleLoginWithUser
// @Summary Login Google
// @Description  Login the user's google, but the function not use with swagger.
// @Tags User
// @Router /api/v1/users/google/callback [get]
func (u *UserController) GoogleLoginWithUser(c *gin.Context) {
	code := c.Query("code")

	userData, response, err := u.UserUseCase.LoginGoogle(c, code)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	c.SetCookie("access_token", response.AccessToken, 0, "/", "localhost", false, true)
	c.SetCookie("refresh_token", response.RefreshToken, 0, "/", "localhost", false, true)
	c.SetCookie("logged_in", response.IsLogged, 0, "/", "localhost", false, false)

	c.JSON(http.StatusOK, gin.H{
		"token": response.SignedToken,
		"user":  userData.User,
	})
}

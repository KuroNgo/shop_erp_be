package user_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	userdomain "shop_erp_mono/domain/human_resource_management/user"
	"shop_erp_mono/pkg/token"
)

// LoginUser
// @Summary Login user
// @Description Login user
// @Tags User
// @Accept json
// @Produce json
// @Param LoginUserRequestDto body user_domain.SignIn true "User data"
// @Security ApiKeyAuth
// @Router /api/v1/users/login [post]
func (l *UserController) LoginUser(ctx *gin.Context) {
	//  Lấy thông tin từ request
	var adminInput userdomain.SignIn
	if err := ctx.ShouldBindJSON(&adminInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error()},
		)
		return
	}

	var userInput userdomain.SignIn
	userInput.Email = adminInput.Email
	userInput.Password = adminInput.Password

	// Tìm kiếm user trong database
	user, err := l.UserUseCase.Login(ctx, userInput)
	if err == nil && user.Verified == true {
		// Generate token
		accessToken, err := token.CreateToken(l.Database.AccessTokenExpiresIn, user.ID, l.Database.AccessTokenPrivateKey)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err.Error()},
			)
			return
		}

		refreshToken, err := token.CreateToken(l.Database.RefreshTokenExpiresIn, user.ID, l.Database.RefreshTokenPrivateKey)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err.Error()},
			)
			return
		}

		ctx.SetCookie("access_token", accessToken, l.Database.AccessTokenMaxAge*1000, "/", "localhost", false, true)
		ctx.SetCookie("refresh_token", refreshToken, l.Database.AccessTokenMaxAge*1000, "/", "localhost", false, true)
		ctx.SetCookie("logged_in", "true", l.Database.AccessTokenMaxAge*1000, "/", "localhost", false, false)

		ctx.JSON(http.StatusOK, gin.H{
			"status":       "success",
			"message":      "Login successful with user role",
			"access_token": accessToken,
		})
		return
	}

	// Trả về thông báo login không thành công
	ctx.JSON(http.StatusBadRequest, gin.H{
		"status":  "error",
		"message": err.Error(),
	})
}

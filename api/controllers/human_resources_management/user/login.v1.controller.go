package user_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	userdomain "shop_erp_mono/domain/human_resource_management/user"
)

// LoginUser
// @Summary Login user
// @Description Login user
// @Tags User
// @Accept json
// @Produce json
// @Param LoginUser body user_domain.SignIn true "User data"
// @Security ApiKeyAuth
// @Router /api/v1/users/login [post]
func (l *UserController) LoginUser(ctx *gin.Context) {
	//  Lấy thông tin từ request
	var userInput userdomain.SignIn
	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error()},
		)
		return
	}

	// Xử lý logic nghiệp vụ và tìm kiếm người dùng
	user, err := l.UserUseCase.LoginUser(ctx, &userInput)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.SetCookie("access_token", user.AccessToken, l.Database.AccessTokenMaxAge*1000, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", user.RefreshToken, l.Database.AccessTokenMaxAge*1000, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", user.IsLogged, l.Database.AccessTokenMaxAge*1000, "/", "localhost", false, false)

	// Trả về phản hồi thành công
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}

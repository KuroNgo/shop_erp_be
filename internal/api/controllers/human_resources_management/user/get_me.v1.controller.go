package user_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	// Lấy cookie access_token từ request
	cookie, err := ctx.Cookie("access_token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "fail",
			"message": "You are not logged in!",
		})
		return
	}

	// Gọi use case để xử lý logic nghiệp vụ
	result, err := u.UserUseCase.GetByIDForCheckCookie(ctx, cookie)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status":  "fail",
			"message": "Failed to get user data: " + err.Error(),
		})
		return
	}

	// Trả về phản hồi thành công
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   result,
	})
}

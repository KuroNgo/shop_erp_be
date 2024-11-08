package user_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	userdomain "shop_erp_mono/internal/domain/human_resource_management/user"
)

// UpdateUser updates the user's information
// @Summary Update User Information
// @Description Updates the user's first name, last name, and username
// @Tags User
// @Accept json
// @Produce json
// @Router /api/v1/users/update [put]
// @Security CookieAuth
func (u *UserController) UpdateUser(ctx *gin.Context) {
	currentUser, exists := ctx.Get("currentUser")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "fail",
			"message": "You are not logged in!",
		})
		return
	}

	fullName := ctx.Request.FormValue("full_name")
	input := userdomain.Input{
		Username: fullName,
	}

	file, _ := ctx.FormFile("file")
	err := u.UserUseCase.UpdateOne(ctx, fmt.Sprint(currentUser), &input, file)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "You are not logged in!",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Updated user",
	})

}

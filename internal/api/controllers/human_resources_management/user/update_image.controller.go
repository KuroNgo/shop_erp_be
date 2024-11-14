package user_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpdateImage updates the user's information
// @Summary Update User Information
// @Description Updates the user's first name, last name, and username
// @Tags User
// @Accept json
// @Produce json
// @Router /api/v1/users/update [put]
// @Security CookieAuth
func (u *UserController) UpdateImage(ctx *gin.Context) {
	currentUser, exists := ctx.Get("currentUser")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "fail",
			"message": "You are not logged in!",
		})
		return
	}

	file, _ := ctx.FormFile("file")
	err := u.UserUseCase.UpdateImage(ctx, fmt.Sprint(currentUser), file)
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

package user_controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	user_domain "shop_erp_mono/domain/human_resource_management/user"
	"shop_erp_mono/pkg/helper"
	"time"
)

// UpdateUser updates the user's information
// @Summary Update User Information
// @Description Updates the user's first name, last name, and username
// @Tags User
// @Accept json
// @Produce json
// @@Success 200 {object} "status: success, message:update user success"
// @Failure 400 {object} map[string]interface{} "status: fail, message: detailed error message"
// @Failure 401 {object} map[string]interface{} "status: fail, message: You are not logged in!"
// @Router /api/v1/users/update [patch]
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

	user, err := u.UserUseCase.GetByID(ctx, fmt.Sprint(currentUser))
	if err != nil || user == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "Unauthorized",
			"message": "You are not authorized to perform this action!",
		})
		return
	}

	fullName := ctx.Request.FormValue("full_name")

	file, err := ctx.FormFile("file")
	if err != nil {
		userResponse := user_domain.UpdateUser{
			ID:        user.ID,
			Username:  fullName,
			UpdatedAt: time.Now(),
		}

		err = u.UserUseCase.Update(ctx, &userResponse)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
		return
	}

	if !helper.IsImage(file.Filename) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid file format. Only images are allowed.",
		})
		return
	}

	// Mở file để đọc dữ liệu
	f, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error opening uploaded file",
		})
		return
	}
	defer func(f multipart.File) {
		err = f.Close()
		if err != nil {
			return
		}
	}(f)

	resultString, err := json.Marshal(user)
	userResponse := user_domain.UpdateUser{
		ID:        user.ID,
		Username:  fullName,
		UpdatedAt: time.Now(),
	}

	err = u.UserUseCase.Update(ctx, &userResponse)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status":  "fail",
			"message": string(resultString) + "the user belonging to this token no logger exists",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Updated user",
	})

}

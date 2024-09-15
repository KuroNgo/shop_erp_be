package user_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteCurrentUser delete the user's information
// @Summary Delete User Information
// @Description Deletes the user's information
// @Tags User
// @Accept json
// @Produce json
// @Router /api/v1/users/delete [delete]
// @Security CookieAuth
func (u *UserController) DeleteCurrentUser(c *gin.Context) {
	currentUser, exist := c.Get("currentUser")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "You are not login!",
		})
		return
	}

	err := u.UserUseCase.DeleteOne(c, fmt.Sprint(currentUser))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

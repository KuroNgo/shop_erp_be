package account_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByName get by name the account's information
// @Summary Get Account Information
// @Description Get the account's information
// @Tags Account
// @Accept json
// @Produce json
// @Param name path string true "Attendance ID"
// @Router /api/v1/accounts/get/name [get]
// @Security CookieAuth
func (a *AccountController) GetByName(ctx *gin.Context) {
	name := ctx.Query("name")

	data, err := a.AccountUseCase.GetByName(ctx, name)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}

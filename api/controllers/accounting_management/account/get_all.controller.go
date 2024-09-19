package account_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAll Get the account's information
// @Summary Get Account Information
// @Description Get the account's information
// @Tags Account
// @Accept json
// @Produce json
// @Router /api/v1/accounts/get/all [get]
// @Security CookieAuth
func (a *AccountController) GetAll(ctx *gin.Context) {
	data, err := a.AccountUseCase.GetAll(ctx)
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

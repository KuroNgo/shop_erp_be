package account_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByID get by id the account's information
// @Summary Get Account Information
// @Description Get the account's information
// @Tags Account
// @Accept json
// @Produce json
// @Param _id path string true "Attendance ID"
// @Router /api/v1/accounts/get/_id [get]
// @Security CookieAuth
func (a *AccountController) GetByID(ctx *gin.Context) {
	_id := ctx.Query("_id")

	data, err := a.AccountUseCase.GetByID(ctx, _id)
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

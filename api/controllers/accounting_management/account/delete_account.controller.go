package account_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteAccount delete the account's information
// @Summary Delete Account Information
// @Description Delete the account's information
// @Tags Account
// @Accept json
// @Produce json
// @Param _id path string true "Account ID"
// @Router /api/v1/accounts/delete [delete]
// @Security CookieAuth
func (a *AccountController) DeleteAccount(ctx *gin.Context) {
	_id := ctx.Param("_id")

	err := a.AccountUseCase.DeleteAccount(ctx, _id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

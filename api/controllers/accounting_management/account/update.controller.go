package account_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	accountdomain "shop_erp_mono/domain/accounting_management/account"
)

// UpdateOne update the account's information
// @Summary Update Account Information
// @Description Update the account's information
// @Tags Account
// @Accept json
// @Produce json
// @Param Account body account_domain.Input true "Account data"
// @Param name path string true "Attendance ID"
// @Router /api/v1/accounts/update [put]
// @Security CookieAuth
func (a *AccountController) UpdateOne(ctx *gin.Context) {
	var input accountdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	_id := ctx.Param("_id")

	if err := a.AccountUseCase.UpdateOne(ctx, _id, &input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

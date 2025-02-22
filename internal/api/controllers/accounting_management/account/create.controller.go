package account_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	accountdomain "shop_erp_mono/internal/domain/accounting_management/account"
)

// CreateOne create the account's information
// @Summary Create Account Information
// @Description Create the account's information
// @Tags Account
// @Accept json
// @Produce json
// @Param Account body account_domain.Input true "Account data"
// @Router /api/v1/accounts/create [post]
// @Security CookieAuth
func (a *AccountController) CreateOne(ctx *gin.Context) {
	var input accountdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	err := a.AccountUseCase.CreateOne(ctx, &input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

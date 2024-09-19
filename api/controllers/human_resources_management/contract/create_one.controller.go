package contract_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	contracts_domain "shop_erp_mono/domain/human_resource_management/contracts"
)

// CreateOne create the contract's information
// @Summary Create Contract Information
// @Description Create the contract's information
// @Tags Contract
// @Accept json
// @Produce json
// @Param Contract body contracts_domain.Input true "Contract data"
// @Router /api/v1/contracts/create [post]
// @Security CookieAuth
func (c *ContractController) CreateOne(ctx *gin.Context) {
	var input contracts_domain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	err := c.ContractUseCase.CreateOne(ctx, &input)
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

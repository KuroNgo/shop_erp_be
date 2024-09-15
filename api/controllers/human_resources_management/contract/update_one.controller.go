package contract_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	contracts_domain "shop_erp_mono/domain/human_resource_management/contracts"
)

// UpdateOneContract create the contract's information
// @Summary Update Contract Information
// @Description Update the contract's information
// @Tags Contract
// @Accept json
// @Produce json
// @Param Contract body contracts_domain.Input true "Contract data"
// @Param _id path string true "Contract ID"
// @Router /api/v1/contracts/update [put]
// @Security CookieAuth
func (c *ContractController) UpdateOneContract(ctx *gin.Context) {
	var input contracts_domain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	id := ctx.Param("_id")

	err := c.ContractUseCase.UpdateOne(ctx, id, &input)
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

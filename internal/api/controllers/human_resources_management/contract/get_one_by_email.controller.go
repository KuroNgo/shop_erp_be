package contract_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByEmail Get one the contract's information
// @Summary Get one Contract Information
// @Description Get one the contract's information by ID
// @Tags Contract
// @Produce json
// @Param email query string true "Contract ID"
// @Router /api/v1/contracts/get/email [get]
// @Security CookieAuth
func (c *ContractController) GetByEmail(ctx *gin.Context) {
	id := ctx.Query("email")

	data, err := c.ContractUseCase.GetByEmail(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
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

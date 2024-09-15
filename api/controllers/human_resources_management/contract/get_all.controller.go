package contract_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// FetchAllContract retrieves the contract's information
// @Summary Get Contract Information
// @Description Retrieves the contract's information
// @Tags Contract
// @Accept  json
// @Produce  json
// @Router /api/v1/contracts/get/all [get]
// @Security CookieAuth
func (c *ContractController) FetchAllContract(ctx *gin.Context) {
	data, err := c.ContractUseCase.GetAll(ctx)
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

package contract_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByID get the contract's information
// @Summary Get Contract Information
// @Description Deletes the contract's information by ID
// @Tags Contract
// @Produce json
// @Param _id query string true "Contract ID"
// @Router /api/v1/contracts/get/_id [get]
// @Security CookieAuth
func (c *ContractController) GetByID(ctx *gin.Context) {
	id := ctx.Query("_id")

	data, err := c.ContractUseCase.GetByID(ctx, id)
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

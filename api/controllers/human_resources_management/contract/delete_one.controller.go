package contract_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOne delete the contract's information
// @Summary Delete Contract Information
// @Description Deletes the contract's information by ID
// @Tags Contract
// @Produce json
// @Param _id query string true "Contract ID"
// @Router /api/v1/contracts/delete/_id [delete]
// @Security CookieAuth
func (c *ContractController) DeleteOne(ctx *gin.Context) {
	id := ctx.Query("_id")

	err := c.ContractUseCase.DeleteOne(ctx, id)
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

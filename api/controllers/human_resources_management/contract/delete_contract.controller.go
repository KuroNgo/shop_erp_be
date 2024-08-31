package contract_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOneContract delete the contract's information
// @Summary Delete Contract Information
// @Description Deletes the contract's information by ID
// @Tags Contract
// @Produce json
// @Param _id path string true "Contract ID"
// @Router /api/v1/contracts/delete/{_id} [delete]
// @Security CookieAuth
func (c *ContractController) DeleteOneContract(ctx *gin.Context) {
	id := ctx.Param("_id")

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

package contract_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// FetchOneByIDContract get the contract's information
// @Summary Get Contract Information
// @Description Deletes the contract's information by ID
// @Tags Contract
// @Produce json
// @Param _id path string true "Contract ID"
// @Router /api/v1/contracts/get/_id [get]
// @Security CookieAuth
func (c *ContractController) FetchOneByIDContract(ctx *gin.Context) {
	id := ctx.Query("_id")

	data, err := c.ContractUseCase.GetOneByID(ctx, id)
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

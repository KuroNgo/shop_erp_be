package contract_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// FetchOneByEmailContract delete the contract's information
// @Summary Delete Contract Information
// @Description Deletes the contract's information by ID
// @Tags Contract
// @Produce json
// @Param _id path string true "Contract ID"
// @Router /api/v1/contracts/get/{email} [get]
// @Security CookieAuth
func (c *ContractController) FetchOneByEmailContract(ctx *gin.Context) {
	id := ctx.Param("email")

	data, err := c.ContractUseCase.GetOneByEmail(ctx, id)
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

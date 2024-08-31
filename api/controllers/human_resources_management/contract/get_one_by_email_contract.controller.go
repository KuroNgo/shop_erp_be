package contract_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// FetchOneByEmailContract Get one the contract's information
// @Summary Get one Contract Information
// @Description Get one the contract's information by ID
// @Tags Contract
// @Produce json
// @Param email path string true "Contract ID"
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

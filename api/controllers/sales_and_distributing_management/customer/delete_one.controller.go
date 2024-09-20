package customer_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOne godoc
// @Summary Delete a customer by ID
// @Description Delete a customer from the system using their ID
// @Tags Customers
// @Accept json
// @Produce json
// @Param _id path string true "Customer ID"
// @Router /api/v1/customers/delete [delete]
func (c *CustomerController) DeleteOne(ctx *gin.Context) {
	_id := ctx.Query("_id")

	if err := c.CustomerUseCase.DeleteOne(ctx, _id); err != nil {
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

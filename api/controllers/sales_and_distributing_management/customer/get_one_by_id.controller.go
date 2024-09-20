package customer_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetOneByID godoc
// @Summary Get a customer by ID
// @Description Retrieve a customer from the system using their ID
// @Tags Customers
// @Accept json
// @Produce json
// @Param _id path string true "Customer ID"
// @Router /api/v1/customers/get/_id [get]
func (c *CustomerController) GetOneByID(ctx *gin.Context) {
	_id := ctx.Query("_id")

	data, err := c.CustomerUseCase.GetOneByID(ctx, _id)
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

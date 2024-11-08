package customer_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAll godoc
// @Summary Get all customers
// @Description Retrieve all customers from the system
// @Tags Customers
// @Accept json
// @Produce json
// @Router /api/v1/customers/get/all [get]
func (c *CustomerController) GetAll(ctx *gin.Context) {
	data, err := c.CustomerUseCase.GetAll(ctx)
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

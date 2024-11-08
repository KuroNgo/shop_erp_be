package customer_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	customerdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/customer"
)

// UpdateOne godoc
// @Summary Update a customer by ID
// @Description Update a customer's information in the system using their ID
// @Tags Customers
// @Accept json
// @Produce json
// @Param _id path string true "Customer ID"
// @Param customer body customer_domain.Input true "Customer data"
// @Router /api/v1/customers/update [put]
func (c *CustomerController) UpdateOne(ctx *gin.Context) {
	var customer customerdomain.Input
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	_id := ctx.Query("_id")

	err := c.CustomerUseCase.UpdateOne(ctx, _id, &customer)
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

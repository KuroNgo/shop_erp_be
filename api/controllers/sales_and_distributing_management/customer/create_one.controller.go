package customer_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	customerdomain "shop_erp_mono/domain/sales_and_distribution_management/customer"
)

// CreateOne godoc
// @Summary Create a new customer
// @Description Create a new customer in the system
// @Tags Customers
// @Accept json
// @Produce json
// @Param customer body customer_domain.Input true "Customer data"
// @Router /api/v1/customers/create [post]
func (c *CustomerController) CreateOne(ctx *gin.Context) {
	var customer customerdomain.Input
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if err := c.CustomerUseCase.CreateOne(ctx, &customer); err != nil {
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

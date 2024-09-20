package sales_order_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByCustomerID godoc
// @Summary Get Sales Orders by Customer ID
// @Description Retrieve all sales orders associated with a specific Customer ID
// @Tags SalesOrders
// @Accept json
// @Produce json
// @Param customer_id query string true "Customer ID"
// @Router /api/v1/sales-orders/get/customer_id [get]
func (s *SalesOrderController) GetByCustomerID(ctx *gin.Context) {
	customerId := ctx.Query("customer_id")

	data, err := s.SalesOrderUseCase.GetByCustomerID(ctx, customerId)
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

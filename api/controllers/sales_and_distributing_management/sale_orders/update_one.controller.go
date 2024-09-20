package sales_order_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	sale_orders_domain "shop_erp_mono/domain/sales_and_distribution_management/sale_orders"
)

// UpdateOne godoc
// @Summary Update a Sales Order
// @Description This API updates an existing sales order based on the provided ID and input data
// @Tags SalesOrders
// @Accept json
// @Produce json
// @Param _id query string true "Sales Order ID"
// @Param salesOrder body sale_orders_domain.Input true "Updated Sales Order information"
// @Router /api/v1/sales-orders/update [put]
func (s *SalesOrderController) UpdateOne(ctx *gin.Context) {
	var input sale_orders_domain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	_id := ctx.Query("_id")

	if err := s.SalesOrderUseCase.UpdateOne(ctx, _id, &input); err != nil {
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

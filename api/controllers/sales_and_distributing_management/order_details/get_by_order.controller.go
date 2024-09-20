package order_detail_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByOrderID godoc
// @Summary Get Order Details by Order ID
// @Description This API retrieves all Order Details associated with a specific Order ID
// @Tags OrderDetails
// @Accept json
// @Produce json
// @Param order_id query string true "Order ID"
// @Router /api/v1/order-details/get/order_id [get]
func (o *OrderDetailController) GetByOrderID(ctx *gin.Context) {
	orderId := ctx.Query("order_id")

	data, err := o.OrderDetailUseCase.GetByOrderID(ctx, orderId)
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

package shipping_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByOrderID godoc
// @Summary Get Shipping Entries by Order ID
// @Description Retrieve shipping entries associated with a specific order ID
// @Tags Shipping
// @Accept json
// @Produce json
// @Param order_id query string true "Order ID"
// @Router /api/v1/shipping/get/order_id [get]
func (s *ShippingController) GetByOrderID(ctx *gin.Context) {
	orderId := ctx.Query("order_id")

	data, err := s.ShippingUseCase.GetByOrderID(ctx, orderId)
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

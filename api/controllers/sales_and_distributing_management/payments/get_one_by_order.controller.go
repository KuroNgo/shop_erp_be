package payment_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByOrder godoc
// @Summary Get Payments by Order ID
// @Description Retrieve all payments associated with a specific Order ID
// @Tags Payments
// @Accept json
// @Produce json
// @Param order_id query string true "Order ID"
// @Router /api/v1/payments/get/order_id [get]
func (p *PaymentController) GetByOrder(ctx *gin.Context) {
	orderId := ctx.Query("order_id")

	data, err := p.PaymentUseCase.GetByStatus(ctx, orderId)
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

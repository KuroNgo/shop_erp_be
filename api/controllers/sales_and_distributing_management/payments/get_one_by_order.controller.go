package payment_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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

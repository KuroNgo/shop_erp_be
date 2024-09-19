package payment_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	payments_domain "shop_erp_mono/domain/sales_and_distribution_management/payments"
)

func (p *PaymentController) UpdateOne(ctx *gin.Context) {
	var input payments_domain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	_id := ctx.Query("_id")

	err := p.PaymentUseCase.UpdateOne(ctx, _id, &input)
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

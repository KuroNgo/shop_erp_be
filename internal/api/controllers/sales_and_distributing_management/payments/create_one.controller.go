package payment_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	payments_domain "shop_erp_mono/internal/domain/sales_and_distribution_management/payments"
)

// CreateOne godoc
// @Summary Create a new Payment
// @Description This API creates a new payment from the input data
// @Tags Payments
// @Accept json
// @Produce json
// @Param payment body payments_domain.Input true "Payment information"
// @Router /api/v1/payments/create [post]
func (p *PaymentController) CreateOne(ctx *gin.Context) {
	var input payments_domain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if err := p.PaymentUseCase.CreateOne(ctx, &input); err != nil {
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

package payment_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByStatus godoc
// @Summary Get Payments by Status
// @Description Retrieve all payments associated with a specific status
// @Tags Payments
// @Accept json
// @Produce json
// @Param status query string true "Payment Status"
// @Router /api/v1/payments/get/status [get]
func (p *PaymentController) GetByStatus(ctx *gin.Context) {
	status := ctx.Query("status")

	data, err := p.PaymentUseCase.GetByStatus(ctx, status)
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

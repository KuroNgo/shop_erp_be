package payment_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByID godoc
// @Summary Get a Payment by ID
// @Description Retrieve a payment based on the provided ID
// @Tags Payments
// @Accept json
// @Produce json
// @Param _id query string true "Payment ID"
// @Router /api/v1/payments/get/_id [get]
func (p *PaymentController) GetByID(ctx *gin.Context) {
	_id := ctx.Query("_id")

	data, err := p.PaymentUseCase.GetByID(ctx, _id)
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

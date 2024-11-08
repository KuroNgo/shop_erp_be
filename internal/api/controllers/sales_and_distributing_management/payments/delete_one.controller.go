package payment_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOne godoc
// @Summary Delete a Payment
// @Description This API deletes a payment based on the provided ID
// @Tags Payments
// @Accept json
// @Produce json
// @Param _id query string true "Payment ID"
// @Router /api/v1/payments/delete [delete]
func (p *PaymentController) DeleteOne(ctx *gin.Context) {
	_id := ctx.Query("_id")

	if err := p.PaymentUseCase.DeleteOne(ctx, _id); err != nil {
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

package payment_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (p *PaymentController) GetByStatus(ctx *gin.Context) {
	status := ctx.Param("status")

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

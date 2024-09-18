package payment_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (p *PaymentController) DeleteOne(ctx *gin.Context) {
	_id := ctx.Param("_id")

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

package order_detail_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (o *OrderDetailController) GetByOrderID(ctx *gin.Context) {
	orderId := ctx.Param("order_id")

	data, err := o.OrderDetailUseCase.GetByOrderID(ctx, orderId)
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
package shipping_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *ShippingController) GetByOrderID(ctx *gin.Context) {
	orderId := ctx.Query("order_id")

	data, err := s.ShippingUseCase.GetByOrderID(ctx, orderId)
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

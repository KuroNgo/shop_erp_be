package order_detail_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (o *OrderDetailController) GetByProductID(ctx *gin.Context) {
	productId := ctx.Param("product_id")

	data, err := o.OrderDetailUseCase.GetByProductID(ctx, productId)
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

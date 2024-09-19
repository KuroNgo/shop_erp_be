package order_detail_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (o *OrderDetailController) GetByID(ctx *gin.Context) {
	_id := ctx.Query("_id")

	data, err := o.OrderDetailUseCase.GetByID(ctx, _id)
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

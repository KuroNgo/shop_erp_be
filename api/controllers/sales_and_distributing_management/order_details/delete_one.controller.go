package order_detail_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (o *OrderDetailController) DeleteOne(ctx *gin.Context) {
	_id := ctx.Query("_id")

	if err := o.OrderDetailUseCase.DeleteOne(ctx, _id); err != nil {
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

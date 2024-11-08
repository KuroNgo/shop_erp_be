package order_detail_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOne godoc
// @Summary Delete an Order Detail
// @Description This API deletes an Order Detail based on the provided ID
// @Tags OrderDetails
// @Accept json
// @Produce json
// @Param _id query string true "Order Detail ID"
// @Router /api/v1/order-details/delete [delete]
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

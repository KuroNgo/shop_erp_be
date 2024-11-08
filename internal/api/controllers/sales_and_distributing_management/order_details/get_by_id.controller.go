package order_detail_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByID godoc
// @Summary Get an Order Detail by ID
// @Description This API retrieves an Order Detail based on the provided ID
// @Tags OrderDetails
// @Accept json
// @Produce json
// @Param _id query string true "Order Detail ID"
// @Router /api/v1/order-details/get/_id [get]
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

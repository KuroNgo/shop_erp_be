package order_detail_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	order_details_domain "shop_erp_mono/internal/domain/sales_and_distribution_management/order_details"
)

// UpdateOne godoc
// @Summary Update an Order Detail
// @Description This API updates an existing Order Detail based on the provided ID and input data
// @Tags OrderDetails
// @Accept json
// @Produce json
// @Param _id query string true "Order Detail ID"
// @Param orderDetail body order_details_domain.Input true "Updated Order Detail information"
// @Router /api/v1/order-details/update [put]
func (o *OrderDetailController) UpdateOne(ctx *gin.Context) {
	var input order_details_domain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	_id := ctx.Query("_id")

	if err := o.OrderDetailUseCase.UpdateOne(ctx, _id, &input); err != nil {
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

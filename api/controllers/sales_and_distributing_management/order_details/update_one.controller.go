package order_detail_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	order_details_domain "shop_erp_mono/domain/sales_and_distribution_management/order_details"
)

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

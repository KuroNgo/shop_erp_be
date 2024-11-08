package order_detail_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	orderdetailsdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/order_details"
)

// CreateOne godoc
// @Summary Create a new Order Detail
// @Description This API creates a new Order Detail from the input data
// @Tags OrderDetails
// @Accept json
// @Produce json
// @Param orderDetail body order_details_domain.Input true "Order Detail information"
// @Router /api/v1/order-details/create [post]
func (o *OrderDetailController) CreateOne(ctx *gin.Context) {
	var input orderdetailsdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if err := o.OrderDetailUseCase.CreateOne(ctx, &input); err != nil {
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

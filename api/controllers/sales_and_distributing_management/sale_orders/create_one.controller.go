package sales_order_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	sale_orders_domain "shop_erp_mono/domain/sales_and_distribution_management/sale_orders"
)

func (s *SalesOrderController) CreateOne(ctx *gin.Context) {
	var input sale_orders_domain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if err := s.SalesOrderUseCase.CreateOne(ctx, &input); err != nil {
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

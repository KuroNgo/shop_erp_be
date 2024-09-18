package shipping_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	shippingdomain "shop_erp_mono/domain/sales_and_distribution_management/shipping"
)

func (s *ShippingController) CreateOne(ctx *gin.Context) {
	var input shippingdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if err := s.ShippingUseCase.CreateOne(ctx, &input); err != nil {
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

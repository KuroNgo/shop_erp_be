package stock_adjustment_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	stock_adjustment_domain "shop_erp_mono/domain/warehouse_management/stock_adjustment"
)

func (s *StockAdjustmentController) UpdateOne(ctx *gin.Context) {
	var input stock_adjustment_domain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	_id := ctx.Param("_id")

	err := s.StockAdjustmentUseCase.UpdateOne(ctx, _id, &input)
	if err != nil {
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

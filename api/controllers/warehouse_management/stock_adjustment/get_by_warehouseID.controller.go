package stock_adjustment_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *StockAdjustmentController) GetByWarehouseID(ctx *gin.Context) {
	warehouseID := ctx.Query("warehouse_id")

	data, err := s.StockAdjustmentUseCase.GetByWarehouseID(ctx, warehouseID)
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

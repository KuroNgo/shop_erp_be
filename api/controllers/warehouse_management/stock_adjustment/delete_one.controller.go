package stock_adjustment_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOne godoc
// @Summary Delete a stock adjustment by ID
// @Description Delete a stock adjustment from the system using its ID
// @Tags StockAdjustment
// @Accept json
// @Produce json
// @Param _id path string true "Stock Adjustment ID"
// @Router /stock-adjustments/delete/{_id} [delete]
func (s *StockAdjustmentController) DeleteOne(ctx *gin.Context) {
	_id := ctx.Query("_id")

	err := s.StockAdjustmentUseCase.DeleteOne(ctx, _id)
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

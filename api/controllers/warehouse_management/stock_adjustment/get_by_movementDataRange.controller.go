package stock_adjustment_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByMovementDataRange godoc
// @Summary Get stock adjustments by date range
// @Description Retrieve stock adjustments within a specified date range
// @Tags stock_adjustments
// @Accept json
// @Produce json
// @Param startDate path string true "Start date in format YYYY-MM-DD"
// @Param endDate path string true "End date in format YYYY-MM-DD"
// @Router /stock-adjustments/get/by-date-range/{startDate}/{endDate} [get]
func (s *StockAdjustmentController) GetByMovementDataRange(ctx *gin.Context) {
	startDateStr := ctx.Query("startDate")
	endDateStr := ctx.Query("endDate")

	data, err := s.StockAdjustmentUseCase.GetByAdjustmentDateRange(ctx, startDateStr, endDateStr)
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

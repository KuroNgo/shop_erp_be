package stock_adjustment_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (s *StockAdjustmentController) GetByMovementDataRange(ctx *gin.Context) {
	startDateStr := ctx.Query("startDate")
	endDateStr := ctx.Query("endDate")

	if startDateStr == "" || endDateStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "startDate and endDate are required",
		})
		return
	}

	// Định dạng thời gian mà bạn mong đợi (ví dụ: "2006-01-02" cho định dạng YYYY-MM-DD)
	layout := "2006-01-02"

	// Chuyển đổi startDate từ chuỗi sang time.Time
	startDate, err := time.Parse(layout, startDateStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid startDate format. Expected format: YYYY-MM-DD",
		})
		return
	}

	// Chuyển đổi endDate từ chuỗi sang time.Time
	endDate, err := time.Parse(layout, endDateStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid endDate format. Expected format: YYYY-MM-DD",
		})
		return
	}

	data, err := s.StockAdjustmentUseCase.GetByAdjustmentDateRange(ctx, startDate, endDate)
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

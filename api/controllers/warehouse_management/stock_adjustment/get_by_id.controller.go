package stock_adjustment_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByID godoc
// @Summary Get a stock adjustment by ID
// @Description Retrieve a stock adjustment from the system using its ID
// @Tags stock_adjustments
// @Accept json
// @Produce json
// @Param _id path string true "Stock Adjustment ID"
// @Router /stock-adjustments/get/_id/{_id} [get]
func (s *StockAdjustmentController) GetByID(ctx *gin.Context) {
	_id := ctx.Query("_id")

	data, err := s.StockAdjustmentUseCase.GetByID(ctx, _id)
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

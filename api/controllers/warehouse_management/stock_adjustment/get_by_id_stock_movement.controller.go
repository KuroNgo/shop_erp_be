package stock_adjustment_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *StockAdjustmentController) GetByID(ctx *gin.Context) {
	_id := ctx.Param("_id")

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

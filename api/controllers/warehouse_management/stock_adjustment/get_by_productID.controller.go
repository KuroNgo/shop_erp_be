package stock_adjustment_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *StockAdjustmentController) GetByProductID(ctx *gin.Context) {
	productId := ctx.Param("product_id")

	data, err := s.StockAdjustmentUseCase.GetByID(ctx, productId)
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

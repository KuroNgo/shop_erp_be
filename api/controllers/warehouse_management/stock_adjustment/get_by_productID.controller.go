package stock_adjustment_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByProductID godoc
// @Summary Get stock adjustments by Product ID
// @Description Retrieve stock adjustments related to a specific Product ID
// @Tags StockAdjustment
// @Accept json
// @Produce json
// @Param product_id path string true "Product ID"
// @Router /stock-adjustments/get/product/{product_id} [get]
func (s *StockAdjustmentController) GetByProductID(ctx *gin.Context) {
	productId := ctx.Query("product_id")

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

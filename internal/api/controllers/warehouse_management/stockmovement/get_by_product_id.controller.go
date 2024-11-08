package stockmovement_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *StockMovementController) GetByProductID(ctx *gin.Context) {
	productID := ctx.Query("product_id")

	data, err := s.StockMovementUseCase.GetByProductID(ctx, productID)
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

package stockmovement_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	stockmovementdomain "shop_erp_mono/internal/domain/warehouse_management/stockmovement"
)

func (s *StockMovementController) CreateOne(ctx *gin.Context) {
	var input stockmovementdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	err := s.StockMovementUseCase.CreateOne(ctx, &input)
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

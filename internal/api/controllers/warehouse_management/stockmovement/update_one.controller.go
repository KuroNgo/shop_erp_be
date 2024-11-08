package stockmovement_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	stockmovement_domain "shop_erp_mono/internal/domain/warehouse_management/stockmovement"
)

func (s *StockMovementController) UpdateOne(ctx *gin.Context) {
	var input stockmovement_domain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	_id := ctx.Query("_id")

	err := s.StockMovementUseCase.UpdateOne(ctx, _id, &input)
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

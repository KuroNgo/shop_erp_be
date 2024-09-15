package stockmovement_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *StockMovementController) DeleteOne(ctx *gin.Context) {
	_id := ctx.Param("_id")

	err := s.StockMovementUseCase.DeleteOne(ctx, _id)
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

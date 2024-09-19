package sales_order_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *SalesOrderController) DeleteOne(ctx *gin.Context) {
	_id := ctx.Query("_id")

	if err := s.SalesOrderUseCase.DeleteOne(ctx, _id); err != nil {
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

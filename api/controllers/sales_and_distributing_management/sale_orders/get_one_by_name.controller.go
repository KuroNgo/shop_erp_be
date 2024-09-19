package sales_order_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *SalesOrderController) GetByStatus(ctx *gin.Context) {
	status := ctx.Query("status")

	data, err := s.SalesOrderUseCase.GetByStatus(ctx, status)
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

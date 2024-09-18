package sales_order_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *SalesOrderController) GetByCustomerID(ctx *gin.Context) {
	customerId := ctx.Param("customer_id")

	data, err := s.SalesOrderUseCase.GetByCustomerID(ctx, customerId)
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

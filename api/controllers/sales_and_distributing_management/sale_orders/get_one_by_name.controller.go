package sales_order_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByStatus godoc
// @Summary Get Sales Orders by Status
// @Description Retrieve all sales orders associated with a specific status
// @Tags SalesOrders
// @Accept json
// @Produce json
// @Param status query string true "Sales Order Status"
// @Router /api/v1/sales-orders/get/status [get]
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

package sales_order_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByID godoc
// @Summary Get a Sales Order by ID
// @Description Retrieve a sales order based on the provided ID
// @Tags SalesOrders
// @Accept json
// @Produce json
// @Param _id query string true "Sales Order ID"
// @Router /api/v1/sales-orders/get/_id [get]
func (s *SalesOrderController) GetByID(ctx *gin.Context) {
	_id := ctx.Query("_id")

	data, err := s.SalesOrderUseCase.GetByID(ctx, _id)
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

package sales_order_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOne godoc
// @Summary Delete a Sales Order
// @Description This API deletes a sales order based on the provided ID
// @Tags SalesOrders
// @Accept json
// @Produce json
// @Param _id query string true "Sales Order ID"
// @Success 200 {object} map[string]interface{} "status: success"
// @Failure 400 {object} map[string]interface{} "status: error, message: Deletion error"
// @Router /api/v1/sales-orders/delete [delete]
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

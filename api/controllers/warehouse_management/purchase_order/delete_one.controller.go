package purchase_order_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOne deletes a purchase order by ID
// @Summary Delete a purchase order
// @Description Delete a purchase order using its ID
// @Tags PurchaseOrder
// @Accept json
// @Produce json
// @Param _id path string true "Purchase Order ID"
// @Router /api/v1/purchase_orders/delete{_id} [delete]
func (p *PurchaseOrderController) DeleteOne(ctx *gin.Context) {
	_id := ctx.Param("_id")

	err := p.PurchaseOrderUseCase.Delete(ctx, _id)
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

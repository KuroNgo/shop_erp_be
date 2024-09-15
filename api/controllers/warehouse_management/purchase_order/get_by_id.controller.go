package purchase_order_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByID retrieves a purchase order by ID
// @Summary Get a purchase order by ID
// @Description Retrieve a purchase order using its ID
// @Tags PurchaseOrder
// @Accept json
// @Produce json
// @Param _id path string true "Purchase Order ID"
// @Router /api/v1/purchase-orders/get/{_id} [get]
func (p *PurchaseOrderController) GetByID(ctx *gin.Context) {
	_id := ctx.Param("_id")

	data, err := p.PurchaseOrderUseCase.GetByID(ctx, _id)
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

package purchase_order_detail_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByIPurchaseOrderD godoc
// @Summary Get purchase order details by Purchase Order ID
// @Description Retrieve purchase order details using the Purchase Order ID
// @Tags PurchaseOrderDetail
// @Accept json
// @Produce json
// @Param purchaseOrder_id path string true "Purchase Order ID"
// @Router /api/v1/purchase_order_details/get/by-order/{purchaseOrder_id} [get]
func (p *PurchaseOrderDetailController) GetByIPurchaseOrderD(ctx *gin.Context) {
	purchaseOrderId := ctx.Query("purchaseOrder_id")

	data, err := p.PurchaseOrderDetailUseCase.GetByPurchaseOrderID(ctx, purchaseOrderId)
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

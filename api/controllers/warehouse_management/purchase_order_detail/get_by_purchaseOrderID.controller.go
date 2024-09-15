package purchase_order_detail_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (p *PurchaseOrderDetailController) GetByIPurchaseOrderD(ctx *gin.Context) {
	purchaseOrderId := ctx.Param("purchaseOrder_id")

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

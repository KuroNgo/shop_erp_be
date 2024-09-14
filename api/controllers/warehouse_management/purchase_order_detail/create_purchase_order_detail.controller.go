package purchase_order_detail_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	purchaseorderdetaildomain "shop_erp_mono/domain/warehouse_management/purchase_order_detail"
)

func (p *PurchaseOrderDetailController) CreateOne(ctx *gin.Context) {
	var input purchaseorderdetaildomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	err := p.PurchaseOrderDetailUseCase.Create(ctx, &input)
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

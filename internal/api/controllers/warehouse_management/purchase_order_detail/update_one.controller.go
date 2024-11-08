package purchase_order_detail_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	purchaseorderdetaildomain "shop_erp_mono/internal/domain/warehouse_management/purchase_order_detail"
)

// UpdateOne godoc
// @Summary Update a purchase order detail
// @Description Update a purchase order detail using its ID
// @Tags PurchaseOrderDetail
// @Accept json
// @Produce json
// @Param _id path string true "Purchase Order Detail ID"
// @Param input body purchase_order_detail_domain.Input true "Purchase Order Detail Input"
// @Router /api/v1/purchase-order-details/update [put]
func (p *PurchaseOrderDetailController) UpdateOne(ctx *gin.Context) {
	var input purchaseorderdetaildomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	_id := ctx.Query("_id")

	err := p.PurchaseOrderDetailUseCase.UpdateOne(ctx, _id, &input)
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

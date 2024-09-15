package purchase_order_detail_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	purchaseorderdetaildomain "shop_erp_mono/domain/warehouse_management/purchase_order_detail"
)

// CreateOne creates a new purchase order detail
// @Summary Create a new purchase order detail
// @Description Create a new detail entry for a purchase order
// @Tags PurchaseOrderDetail
// @Accept json
// @Produce json
// @Param purchase_order_detail body purchase_order_detail_domain.Input true "Purchase Order Detail Input"
// @Router /api/v1/purchase_order_details/create [post]
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

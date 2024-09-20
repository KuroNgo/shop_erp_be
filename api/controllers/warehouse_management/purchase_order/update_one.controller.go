package purchase_order_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	purchaseorderdomain "shop_erp_mono/domain/warehouse_management/purchase_order"
)

// UpdateOne updates a purchase order by ID
// @Summary Update a purchase order by ID
// @Description Update the details of a purchase order using its ID
// @Tags PurchaseOrder
// @Accept json
// @Produce json
// @Param _id path string true "Purchase Order ID"
// @Param purchase_order body purchase_order_domain.Input true "Purchase Order Input"
// @Router /api/v1/purchase-orders/update [put]
func (p *PurchaseOrderController) UpdateOne(ctx *gin.Context) {
	var input purchaseorderdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	_id := ctx.Query("_id")

	err := p.PurchaseOrderUseCase.UpdateOne(ctx, _id, &input)
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

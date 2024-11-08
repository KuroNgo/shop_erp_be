package purchase_order_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	purchaseorderdomain "shop_erp_mono/internal/domain/warehouse_management/purchase_order"
)

// CreateOne creates a new purchase order
// @Summary Create a new purchase order
// @Description Create a new purchase order with the given details
// @Tags PurchaseOrder
// @Accept json
// @Produce json
// @Param input body purchase_order_domain.Input true "Purchase Order Input"
// @Router /api/v1/purchase_orders/create [post]
func (p *PurchaseOrderController) CreateOne(ctx *gin.Context) {
	var input purchaseorderdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	err := p.PurchaseOrderUseCase.CreateOne(ctx, &input)
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

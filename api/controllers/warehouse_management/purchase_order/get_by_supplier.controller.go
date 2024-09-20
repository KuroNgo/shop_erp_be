package purchase_order_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetBySupplierID retrieves purchase orders by supplier ID
// @Summary Get purchase orders by supplier ID
// @Description Retrieve a list of purchase orders associated with a specific supplier ID
// @Tags PurchaseOrder
// @Accept json
// @Produce json
// @Param supplier_id path string true "Supplier ID"
// @Router /api/v1/purchase-orders/get/supplier_id [get]
func (p *PurchaseOrderController) GetBySupplierID(ctx *gin.Context) {
	supplierId := ctx.Query("supplier_id")

	data, err := p.PurchaseOrderUseCase.GetByID(ctx, supplierId)
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

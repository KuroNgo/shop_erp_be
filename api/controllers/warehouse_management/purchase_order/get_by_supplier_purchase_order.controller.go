package purchase_order_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (p *PurchaseOrderController) GetBySupplierID(ctx *gin.Context) {
	supplierId := ctx.Param("supplier_id")

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

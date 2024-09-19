package purchase_order_detail_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAll godoc
// @Summary Get all purchase order details
// @Description Retrieve all purchase order details from the system
// @Tags PurchaseOrderDetail
// @Accept json
// @Produce json
// @Router /api/v1/purchase_order_details/get/all [get]
func (p *PurchaseOrderDetailController) GetAll(ctx *gin.Context) {
	data, err := p.PurchaseOrderDetailUseCase.GetAll(ctx)
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

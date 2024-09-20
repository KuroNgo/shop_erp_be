package purchase_order_detail_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByID godoc
// @Summary Get a purchase order detail by ID
// @Description Retrieve a purchase order detail from the system using its ID
// @Tags PurchaseOrderDetail
// @Accept json
// @Produce json
// @Param _id path string true "Purchase Order Detail ID"
// @Router /api/v1/purchase-order-details/get/_id [get]
func (p *PurchaseOrderDetailController) GetByID(ctx *gin.Context) {
	_id := ctx.Query("_id")

	data, err := p.PurchaseOrderDetailUseCase.GetByID(ctx, _id)
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

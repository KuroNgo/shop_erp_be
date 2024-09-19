package purchase_order_detail_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOne godoc
// @Summary Delete a purchase order detail by ID
// @Description Delete a purchase order detail from the system using its ID
// @Tags PurchaseOrderDetail
// @Accept json
// @Produce json
// @Param _id path string true "Purchase Order Detail ID"
// @Router /api/v1/purchase_order_details/delete/{_id} [delete]
func (p *PurchaseOrderDetailController) DeleteOne(ctx *gin.Context) {
	_id := ctx.Query("_id")

	err := p.PurchaseOrderDetailUseCase.Delete(ctx, _id)
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

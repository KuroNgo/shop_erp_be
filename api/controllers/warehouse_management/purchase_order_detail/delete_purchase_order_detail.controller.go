package purchase_order_detail_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (p *PurchaseOrderDetailController) DeleteOne(ctx *gin.Context) {
	_id := ctx.Param("_id")

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

package purchase_order_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (p *PurchaseOrderController) GetByID(ctx *gin.Context) {
	_id := ctx.Param("_id")

	data, err := p.PurchaseOrderUseCase.GetByID(ctx, _id)
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

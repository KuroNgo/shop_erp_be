package purchase_order_detail_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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
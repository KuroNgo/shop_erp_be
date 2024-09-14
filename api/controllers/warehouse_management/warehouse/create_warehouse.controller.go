package warehouse_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	warehousedomain "shop_erp_mono/domain/warehouse_management/warehouse"
)

func (w *WarehouseController) CreateOne(ctx *gin.Context) {
	var input warehousedomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	err := w.WarehouseUseCase.CreateWarehouse(ctx, &input)
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

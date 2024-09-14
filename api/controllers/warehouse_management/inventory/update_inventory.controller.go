package inventory_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	inventorydomain "shop_erp_mono/domain/warehouse_management/inventory"
)

func (i *InventoryController) Update(ctx *gin.Context) {
	var input inventorydomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	_id := ctx.Param("_id")

	err := i.InventoryUseCase.UpdateInventory(ctx, _id, &input)
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

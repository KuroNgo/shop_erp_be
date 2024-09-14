package inventory_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (i *InventoryController) CheckInventoryAvailability(ctx *gin.Context) {
	warehouseId := ctx.Param("warehouse_id")
	productId := ctx.Param("product_id")
	requiredQuantity := ctx.Param("required_quantity")
	requireQuan, err := strconv.Atoi(requiredQuantity)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	data, err := i.InventoryUseCase.CheckInventoryAvailability(ctx, productId, warehouseId, requireQuan)
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

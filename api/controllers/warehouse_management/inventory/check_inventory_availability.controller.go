package inventory_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CheckInventoryAvailability checks if the required quantity of a product is available in a warehouse
// @Summary Check inventory availability
// @Description Check if the specified quantity of a product is available in a given warehouse
// @Tags Inventory
// @Accept json
// @Produce json
// @Param warehouse_id path string true "Warehouse ID"
// @Param product_id path string true "Product ID"
// @Param required_quantity path string true "Required quantity"
// @Router /api/v1/inventory/check/{warehouse_id}/{product_id}/{required_quantity} [get]
func (i *InventoryController) CheckInventoryAvailability(ctx *gin.Context) {
	warehouseId := ctx.Query("warehouse_id")
	productId := ctx.Query("product_id")
	requiredQuantity := ctx.Query("required_quantity")
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

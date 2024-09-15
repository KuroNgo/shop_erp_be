package inventory_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByWarehouseID retrieves inventory details by warehouse ID
// @Summary Get inventory by warehouse ID
// @Description Retrieve inventory details using the warehouse ID
// @Tags Inventory
// @Accept json
// @Produce json
// @Param warehouse_id path string true "Warehouse ID"
// @Router /api/v1/inventory/warehouse/{warehouse_id} [get]
func (i *InventoryController) GetByWarehouseID(ctx *gin.Context) {
	warehouseId := ctx.Param("warehouse_id")

	data, err := i.InventoryUseCase.GetInventoryByWarehouse(ctx, warehouseId)
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

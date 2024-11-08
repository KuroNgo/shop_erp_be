package inventory_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	inventorydomain "shop_erp_mono/internal/domain/warehouse_management/inventory"
)

// Update updates an inventory item
// @Summary Update inventory by ID
// @Description Update an inventory item using its ID
// @Tags Inventory
// @Accept json
// @Produce json
// @Param _id path string true "Inventory ID"
// @Param input body inventory_domain.Input true "Inventory Input"
// @Router /api/v1/inventory/{_id} [put]
func (i *InventoryController) Update(ctx *gin.Context) {
	var input inventorydomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	_id := ctx.Query("_id")

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

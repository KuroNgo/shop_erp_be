package inventory_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteInventory deletes an inventory record by ID
// @Summary Delete inventory
// @Description Delete an existing inventory record by its ID
// @Tags Inventory
// @Accept json
// @Produce json
// @Param _id path string true "Inventory ID"
// @Router /api/v1/inventory/delete/{_id} [delete]
func (i *InventoryController) DeleteInventory(ctx *gin.Context) {
	_id := ctx.Query("_id")

	err := i.InventoryUseCase.DeleteInventory(ctx, _id)
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

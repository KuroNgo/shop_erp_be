package inventory_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllInventory retrieves all inventory records
// @Summary Get all inventories
// @Description Retrieve all inventory records from the system
// @Tags Inventory
// @Accept json
// @Produce json
// @Router /api/v1/inventory/all [get]
func (i *InventoryController) GetAllInventory(ctx *gin.Context) {
	data, err := i.InventoryUseCase.ListAllInventories(ctx)
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

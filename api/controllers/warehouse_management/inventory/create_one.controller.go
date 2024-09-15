package inventory_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	inventorydomain "shop_erp_mono/domain/warehouse_management/inventory"
)

// CreateInventory creates a new inventory record
// @Summary Create new inventory
// @Description Add a new inventory record to the system
// @Tags Inventory
// @Accept json
// @Produce json
// @Param input body inventorydomain.Input true "Inventory input data"
// @Router /api/v1/inventory/create [post]
func (i *InventoryController) CreateInventory(ctx *gin.Context) {
	var input inventorydomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	err := i.InventoryUseCase.CreateInventory(ctx, &input)
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

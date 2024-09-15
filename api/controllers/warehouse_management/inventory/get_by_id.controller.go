package inventory_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByIDInventory retrieves a single inventory record by ID
// @Summary Get inventory by ID
// @Description Retrieve an inventory record using its ID
// @Tags Inventory
// @Accept json
// @Produce json
// @Param _id path string true "Inventory ID"
// @Router /api/v1/inventory/{_id} [get]
func (i *InventoryController) GetByIDInventory(ctx *gin.Context) {
	_id := ctx.Param("_id")

	data, err := i.InventoryUseCase.GetInventoryByID(ctx, _id)
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

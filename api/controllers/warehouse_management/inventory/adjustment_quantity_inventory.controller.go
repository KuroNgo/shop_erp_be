package inventory_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AdjustmentQuantity update the inventory's information
// @Summary Create Inventory Information
// @Description Create the inventory's information
// @Tags Inventory
// @Accept json
// @Produce json
// @Param _id path string true "Inventory ID"
// @Param adjustment path string true "Inventory adjustment"
// @Router /api/v1/accounts/update/adjustment [put]
// @Security CookieAuth
func (i *InventoryController) AdjustmentQuantity(ctx *gin.Context) {
	_id := ctx.Param("_id")
	adjustment := ctx.Param("adjustment")
	adjusts, err := strconv.Atoi(adjustment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	data, err := i.InventoryUseCase.AdjustInventoryQuantity(ctx, _id, adjusts)
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

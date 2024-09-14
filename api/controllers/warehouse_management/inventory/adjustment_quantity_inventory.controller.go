package inventory_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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

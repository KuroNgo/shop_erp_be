package inventory_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (i *InventoryController) DeleteInventory(ctx *gin.Context) {
	_id := ctx.Param("_id")

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

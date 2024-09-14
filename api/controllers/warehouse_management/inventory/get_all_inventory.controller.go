package inventory_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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

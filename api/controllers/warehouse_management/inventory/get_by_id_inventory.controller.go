package inventory_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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

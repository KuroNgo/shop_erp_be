package warehouse_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (w *WarehouseController) DeleteOne(ctx *gin.Context) {
	_id := ctx.Param("_id")

	err := w.WarehouseUseCase.DeleteWarehouse(ctx, _id)
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

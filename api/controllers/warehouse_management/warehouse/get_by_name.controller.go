package warehouse_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (w *WarehouseController) GetByName(ctx *gin.Context) {
	name := ctx.Query("name")

	data, err := w.WarehouseUseCase.GetByName(ctx, name)
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

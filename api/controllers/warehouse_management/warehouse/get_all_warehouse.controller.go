package warehouse_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (w *WarehouseController) GetAll(ctx *gin.Context) {
	data, err := w.WarehouseUseCase.GetAllWarehouses(ctx)
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

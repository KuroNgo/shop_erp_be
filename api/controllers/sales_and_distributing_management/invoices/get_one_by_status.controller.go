package invoice_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (i *InvoiceController) GetByStatus(ctx *gin.Context) {
	status := ctx.Param("status")
	data, err := i.InvoiceUseCase.GetByStatus(ctx, status)
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

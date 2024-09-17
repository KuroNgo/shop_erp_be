package invoice_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (i *InvoiceController) GetByName(ctx *gin.Context) {
	name := ctx.Param("name")

	data, err := i.InvoiceUseCase.GetByStatus(ctx, name)
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

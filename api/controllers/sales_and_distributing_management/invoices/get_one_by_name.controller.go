package invoice_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByName godoc
// @Summary Get an invoice by name
// @Description Retrieve an invoice from the system using its name
// @Tags Invoices
// @Accept json
// @Produce json
// @Param name path string true "Invoice Name"
// @Router /api/v1/invoices/get/name [get]
func (i *InvoiceController) GetByName(ctx *gin.Context) {
	name := ctx.Query("name")

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

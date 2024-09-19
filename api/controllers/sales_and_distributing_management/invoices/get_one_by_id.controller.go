package invoice_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByID godoc
// @Summary Get an invoice by ID
// @Description Retrieve an invoice from the system using its ID
// @Tags Invoices
// @Accept json
// @Produce json
// @Param _id path string true "Invoice ID"
// @Router /invoices/get/{_id} [get]
func (i *InvoiceController) GetByID(ctx *gin.Context) {
	_id := ctx.Query("_id")

	data, err := i.InvoiceUseCase.GetByStatus(ctx, _id)
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

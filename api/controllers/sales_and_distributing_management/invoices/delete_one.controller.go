package invoice_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOne godoc
// @Summary Delete an invoice by ID
// @Description Delete an invoice from the system using its ID
// @Tags Invoices
// @Accept json
// @Produce json
// @Param _id path string true "Invoice ID"
// @Router /invoices/{_id} [delete]
func (i *InvoiceController) DeleteOne(ctx *gin.Context) {
	_id := ctx.Query("_id")

	if err := i.InvoiceUseCase.DeleteOne(ctx, _id); err != nil {
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

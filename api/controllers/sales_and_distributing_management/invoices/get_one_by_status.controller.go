package invoice_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByStatus godoc
// @Summary Get invoices by status
// @Description Retrieve invoices from the system using their status
// @Tags Invoices
// @Accept json
// @Produce json
// @Param status path string true "Invoice Status"
// @Router /api/v1/invoices/status [get]
func (i *InvoiceController) GetByStatus(ctx *gin.Context) {
	status := ctx.Query("status")

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

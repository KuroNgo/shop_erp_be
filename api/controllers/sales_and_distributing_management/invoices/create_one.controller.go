package invoice_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	invoicesdomain "shop_erp_mono/domain/sales_and_distribution_management/invoices"
)

// CreateOne godoc
// @Summary Create a new invoice
// @Description Create a new invoice in the system
// @Tags Invoices
// @Accept json
// @Produce json
// @Param invoice body invoice_domain.Input true "Invoice data"
// @Router /invoices/create [post]
func (i *InvoiceController) CreateOne(ctx *gin.Context) {
	var invoice invoicesdomain.Input
	if err := ctx.ShouldBindJSON(&invoice); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if err := i.InvoiceUseCase.CreateOne(ctx, &invoice); err != nil {
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

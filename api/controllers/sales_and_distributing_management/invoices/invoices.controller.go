package invoice_controller

import (
	"shop_erp_mono/bootstrap"
	invoices_domain "shop_erp_mono/domain/sales_and_distribution_management/invoices"
)

type InvoiceController struct {
	Database       *bootstrap.Database
	InvoiceUseCase invoices_domain.InvoiceUseCase
}

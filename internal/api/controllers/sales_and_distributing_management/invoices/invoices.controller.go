package invoice_controller

import (
	"shop_erp_mono/internal/config"
	invoices_domain "shop_erp_mono/internal/domain/sales_and_distribution_management/invoices"
)

type InvoiceController struct {
	Database       *config.Database
	InvoiceUseCase invoices_domain.InvoiceUseCase
}

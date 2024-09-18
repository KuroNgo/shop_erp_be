package shipping_controller

import (
	"shop_erp_mono/bootstrap"
	salereportsdomain "shop_erp_mono/domain/sales_and_distribution_management/sale_reports"
)

type SalesReportController struct {
	Database           *bootstrap.Database
	SalesReportUseCase salereportsdomain.ISalesReportUseCase
}

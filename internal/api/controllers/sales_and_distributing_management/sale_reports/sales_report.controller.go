package sales_report_controller

import (
	"shop_erp_mono/internal/config"
	salereportsdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/sale_reports"
)

type SalesReportController struct {
	Database           *config.Database
	SalesReportUseCase salereportsdomain.ISalesReportUseCase
}

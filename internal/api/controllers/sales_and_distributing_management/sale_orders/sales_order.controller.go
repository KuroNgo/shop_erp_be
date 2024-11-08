package sales_order_controller

import (
	"shop_erp_mono/internal/config"
	saleordersdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/sale_orders"
)

type SalesOrderController struct {
	Database          *config.Database
	SalesOrderUseCase saleordersdomain.ISalesOrderUseCase
}

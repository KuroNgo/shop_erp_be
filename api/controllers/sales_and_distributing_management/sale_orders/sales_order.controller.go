package sales_order_controller

import (
	"shop_erp_mono/bootstrap"
	saleordersdomain "shop_erp_mono/domain/sales_and_distribution_management/sale_orders"
)

type SalesOrderController struct {
	Database          *bootstrap.Database
	SalesOrderUseCase saleordersdomain.ISalesOrderUseCase
}

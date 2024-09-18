package order_detail_controller

import (
	"shop_erp_mono/bootstrap"
	order_details_domain "shop_erp_mono/domain/sales_and_distribution_management/order_details"
)

type OrderDetailController struct {
	Database           *bootstrap.Database
	OrderDetailUseCase order_details_domain.IOrderDetailUseCase
}

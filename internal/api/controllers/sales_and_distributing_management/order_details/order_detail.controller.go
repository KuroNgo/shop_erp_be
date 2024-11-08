package order_detail_controller

import (
	"shop_erp_mono/internal/config"
	order_details_domain "shop_erp_mono/internal/domain/sales_and_distribution_management/order_details"
)

type OrderDetailController struct {
	Database           *config.Database
	OrderDetailUseCase order_details_domain.IOrderDetailUseCase
}

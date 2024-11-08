package shipping_controller

import (
	"shop_erp_mono/internal/config"
	shippingdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/shipping"
)

type ShippingController struct {
	Database        *config.Database
	ShippingUseCase shippingdomain.IShippingUseCase
}

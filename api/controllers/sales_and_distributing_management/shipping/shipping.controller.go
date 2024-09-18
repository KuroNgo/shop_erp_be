package shipping_controller

import (
	"shop_erp_mono/bootstrap"
	shippingdomain "shop_erp_mono/domain/sales_and_distribution_management/shipping"
)

type ShippingController struct {
	Database        *bootstrap.Database
	ShippingUseCase shippingdomain.IShippingUseCase
}

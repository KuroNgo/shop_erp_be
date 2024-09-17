package customer_controller

import (
	"shop_erp_mono/bootstrap"
	customerdomain "shop_erp_mono/domain/sales_and_distribution_management/customer"
)

type CustomerController struct {
	Database        *bootstrap.Database
	CustomerUseCase customerdomain.ICustomerUseCase
}

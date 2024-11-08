package customer_controller

import (
	"shop_erp_mono/internal/config"
	customerdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/customer"
)

type CustomerController struct {
	Database        *config.Database
	CustomerUseCase customerdomain.ICustomerUseCase
}

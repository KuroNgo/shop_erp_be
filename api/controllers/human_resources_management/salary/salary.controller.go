package salary_controller

import (
	"shop_erp_mono/bootstrap"
	salary_domain "shop_erp_mono/domain/human_resource_management/salary"
)

type SalaryController struct {
	Database      *bootstrap.Database
	SalaryUseCase salary_domain.ISalaryUseCase
}

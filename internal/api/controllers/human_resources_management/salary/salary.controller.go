package salary_controller

import (
	"shop_erp_mono/internal/config"
	salary_domain "shop_erp_mono/internal/domain/human_resource_management/salary"
)

type SalaryController struct {
	Database      *config.Database
	SalaryUseCase salary_domain.ISalaryUseCase
}

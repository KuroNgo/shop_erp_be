package salary_base

import (
	"shop_erp_mono/internal/config"
	basesalarydomain "shop_erp_mono/internal/domain/human_resource_management/salary_base"
)

type BaseSalaryController struct {
	Database          *config.Database
	BaseSalaryUseCase basesalarydomain.ISalaryUseCase
}

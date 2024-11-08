package employee_controller

import (
	"shop_erp_mono/internal/config"
	employeesdomain "shop_erp_mono/internal/domain/human_resource_management/employees"
)

type EmployeeController struct {
	Database        *config.Database
	EmployeeUseCase employeesdomain.IEmployeeUseCase
}

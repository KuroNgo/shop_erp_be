package employee_controller

import (
	"shop_erp_mono/bootstrap"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
)

type EmployeeController struct {
	Database        *bootstrap.Database
	EmployeeUseCase employeesdomain.IEmployeeUseCase
}

package department_controller

import (
	"shop_erp_mono/bootstrap"
	departmentsdomain "shop_erp_mono/domain/human_resource_management/departments"
)

type DepartmentController struct {
	DepartmentUseCase departmentsdomain.IDepartmentUseCase
	Database          *bootstrap.Database
}

package department_controller

import (
	"shop_erp_mono/internal/config"
	departmentsdomain "shop_erp_mono/internal/domain/human_resource_management/departments"
)

type DepartmentController struct {
	DepartmentUseCase departmentsdomain.IDepartmentUseCase
	Database          *config.Database
}

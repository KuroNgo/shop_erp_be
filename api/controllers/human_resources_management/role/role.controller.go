package role_controller

import (
	"shop_erp_mono/bootstrap"
	roledomain "shop_erp_mono/domain/human_resource_management/role"
)

type RoleController struct {
	Database    *bootstrap.Database
	RoleUseCase roledomain.IRoleUseCase
}

package role_controller

import (
	"shop_erp_mono/internal/config"
	roledomain "shop_erp_mono/internal/domain/human_resource_management/role"
)

type RoleController struct {
	Database    *config.Database
	RoleUseCase roledomain.IRoleUseCase
}

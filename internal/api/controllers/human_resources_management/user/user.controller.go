package user_controller

import (
	"shop_erp_mono/internal/config"
	userdomain "shop_erp_mono/internal/domain/human_resource_management/user"
)

type UserController struct {
	Database    *config.Database
	UserUseCase userdomain.IUserUseCase
}

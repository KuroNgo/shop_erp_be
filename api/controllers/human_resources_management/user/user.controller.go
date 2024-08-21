package user_controller

import (
	"shop_erp_mono/bootstrap"
	userdomain "shop_erp_mono/domain/human_resource_management/user"
)

type UserController struct {
	Database    *bootstrap.Database
	UserUseCase userdomain.IUserUseCase
}

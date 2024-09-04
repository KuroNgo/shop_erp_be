package account_controller

import (
	"shop_erp_mono/bootstrap"
	accountdomain "shop_erp_mono/domain/accounting_management/account"
)

type AccountController struct {
	Database       *bootstrap.Database
	AccountUseCase accountdomain.IAccountUseCase
}

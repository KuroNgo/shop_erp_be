package account_controller

import (
	"shop_erp_mono/internal/config"
	accountdomain "shop_erp_mono/internal/domain/accounting_management/account"
)

type AccountController struct {
	Database       *config.Database
	AccountUseCase accountdomain.IAccountUseCase
}

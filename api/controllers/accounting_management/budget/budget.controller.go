package budget_controller

import (
	"shop_erp_mono/bootstrap"
	budgetsdomain "shop_erp_mono/domain/accounting_management/budgets"
)

type BudgetController struct {
	Database      *bootstrap.Database
	BudgetUseCase budgetsdomain.IBudgetUseCase
}

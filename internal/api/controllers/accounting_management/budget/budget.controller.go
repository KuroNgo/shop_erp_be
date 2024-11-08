package budget_controller

import (
	"shop_erp_mono/internal/config"
	budgetsdomain "shop_erp_mono/internal/domain/accounting_management/budgets"
)

type BudgetController struct {
	Database      *config.Database
	BudgetUseCase budgetsdomain.IBudgetUseCase
}

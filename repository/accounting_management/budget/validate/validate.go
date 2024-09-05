package validate

import (
	"errors"
	budgets_domain "shop_erp_mono/domain/accounting_management/budgets"
)

func IsNilBudget(input *budgets_domain.Input) error {
	if input.BudgetName == "" {
		return errors.New("the budget's information do not nil")
	}

	if input.StartDate.IsZero() {
		return errors.New("the budget's information do not nil")
	}

	if input.EndDate.IsZero() {
		return errors.New("the budget's information do not nil")
	}

	if input.Amount < 0 {
		return errors.New("the budget's data is invalid")
	}

	return nil
}

package validate

import (
	"errors"
	salarydomain "shop_erp_mono/domain/human_resource_management/salary"
)

func IsNilSalary(salary *salarydomain.Input) error {
	if salary.BaseSalary < 0 {
		return errors.New("the salary's information is invalid")
	}

	if salary.Deductions < 0 {
		return errors.New("the salary's information is invalid")
	}

	validCurrencies := []string{"USD", "EUR", "VND"}
	if !contains(validCurrencies, salary.UnitCurrency) {
		return errors.New("invalid currency unit")
	}

	return nil
}

func contains(slice []string, item string) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}

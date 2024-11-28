package validate

import (
	"errors"
	base_salary_domain "shop_erp_mono/internal/domain/human_resource_management/salary_base"
)

func BaseSalary(baseSalary *base_salary_domain.Input) error {
	if baseSalary.BaseSalary < 0 {
		return errors.New("the salary's information is invalid")
	}

	validCurrencies := []string{"USD", "EUR", "VND"}
	if !contains(validCurrencies, baseSalary.UnitCurrency) {
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

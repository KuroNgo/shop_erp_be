package validate

import (
	"errors"
	salarydomain "shop_erp_mono/domain/human_resource_management/salary"
)

func IsNilSalary(salary *salarydomain.Input) error {
	if salary.BaseSalary == 0 {
		return errors.New("the salary's information do not nil")
	}

	if salary.UnitCurrency == "" {
		return errors.New("the salary's information do not nil")
	}

	if salary.Deductions == 0 {
		return errors.New("the salary's information do not nil")
	}

	if salary.NetSalary == 0 {
		return errors.New("the salary's information do not nil")
	}

	return nil
}

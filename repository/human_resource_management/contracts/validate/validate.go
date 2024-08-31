package validate

import (
	"errors"
	contracts_domain "shop_erp_mono/domain/human_resource_management/contracts"
)

func IsNilContract(input *contracts_domain.Input) error {
	if input.ContractType == "" {
		return errors.New("contract type do not nil")
	}

	if input.EmployeeEmail == "" {
		return errors.New("employee email do not nil")
	}

	if input.Salary == 0 {
		return errors.New("salary do not not equal or lower with 0")
	}

	return nil
}

func IsNilEmail(email string) error {
	if email == "" {
		return errors.New("email do not not nil")
	}

	return nil
}

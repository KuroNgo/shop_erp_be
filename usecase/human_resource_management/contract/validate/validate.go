package validate

import (
	"errors"
	contracts_domain "shop_erp_mono/domain/human_resource_management/contracts"
)

func Contract(input *contracts_domain.Input) error {
	if input.ContractType == "" {
		return errors.New("contract type do not nil")
	}

	if input.EmployeeEmail == "" {
		return errors.New("employee email do not nil")
	}

	if input.Salary == 0 {
		return errors.New("salary do not not equal or lower with 0")
	}

	if input.StartDate.Before(input.EndDate) {
		return errors.New("endDate do not before startDate")
	}

	return nil
}

func IsNilEmail(email string) error {
	if email == "" {
		return errors.New("email do not not nil")
	}

	return nil
}

package validate

import (
	"errors"
	employees_domain "shop_erp_mono/domain/human_resource_management/employees"
)

func IsNilEmployee(employee *employees_domain.Input) error {
	if employee.FirstName == "" {
		return errors.New("the employee's information do not nil")
	}

	if employee.LastName == "" {
		return errors.New("the employee's information do not nil")
	}

	if employee.Gender == "" {
		return errors.New("the employee's information do not nil")
	}

	if employee.Email == "" {
		return errors.New("the employee's information do not nil")
	}

	if employee.Phone == "" {
		return errors.New("the employee's information do not nil")
	}

	if employee.Address == "" {
		return errors.New("the employee's information do not nil")
	}

	if employee.AvatarURL == "" {
		return errors.New("the employee's information do not nil")
	}

	if employee.Department == "" {
		return errors.New("the employee's information do not nil")
	}

	if employee.Role == "" {
		return errors.New("the employee's information do not nil")
	}

	return nil
}

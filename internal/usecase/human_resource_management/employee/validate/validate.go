package validate

import (
	"errors"
	employeesdomain "shop_erp_mono/internal/domain/human_resource_management/employees"
	helper2 "shop_erp_mono/pkg/shared/helper"
	"strings"
)

func Employee(employee *employeesdomain.Input) error {
	if employee.FirstName == "" {
		return errors.New("the employee's information do not nil")
	}

	if employee.LastName == "" {
		return errors.New("the employee's information do not nil")
	}

	genderMap := map[string]bool{
		"male":   true,
		"female": true,
		"nam":    true,
		"nữ":     true,
	}

	if !genderMap[strings.ToLower(employee.Gender)] {
		return errors.New("the employee's information is not valid")
	}

	if employee.Email == "" {
		return errors.New("the employee's information do not nil")
	}

	if !helper2.EmailValid(employee.Email) {
		return errors.New("the employee's information do not valid")
	}

	if employee.Phone == "" {
		return errors.New("the employee's information do not nil")
	}

	if !helper2.PhoneValid(employee.Phone) {
		return errors.New("the employee's information do not valid")
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

func IsNilEmailEmployee(email string) error {
	if email == "" {
		return errors.New("the name do not null")
	}
	return nil
}

package validate

import (
	"errors"
	departmentsdomain "shop_erp_mono/domain/human_resource_management/departments"
)

func IsNilDepartment(department *departmentsdomain.Input) error {
	if department.Name == "" {
		return errors.New("the department's information do not null")
	}

	if department.Description == "" {
		return errors.New("the department's information do not null")
	}
	return nil
}

func IsNilDepartment2(department *departmentsdomain.Department) error {
	if department.Name == "" {
		return errors.New("the department's information do not null")
	}

	if department.Description == "" {
		return errors.New("the department's information do not null")
	}
	return nil
}

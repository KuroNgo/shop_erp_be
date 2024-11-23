package validate

import (
	"errors"
	departmentsdomain "shop_erp_mono/internal/domain/human_resource_management/departments"
	"shop_erp_mono/pkg/shared/helper"
)

func Department(department *departmentsdomain.Input) error {
	if helper.AlphabetOnlyRegex(department.Name) {
		return errors.New("the department's information do not null")
	}

	if department.Description == "" {
		return errors.New("the department's information do not null")
	}

	if department.ManagerEmail == "" {
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

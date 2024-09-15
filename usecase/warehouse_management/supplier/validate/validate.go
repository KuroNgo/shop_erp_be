package validate

import (
	"errors"
	supplierdomain "shop_erp_mono/domain/warehouse_management/supplier"
)

func ValidateSupplier(input *supplierdomain.Input) error {
	if input.Address == "" {
		return errors.New("the supplier's information do not nil")
	}

	if input.SupplierName == "" {
		return errors.New("the supplier's information do not nil")
	}

	if input.ContactPerson == "" {
		return errors.New("the supplier's information do not nil")
	}

	if input.PhoneNumber == "" {
		return errors.New("the supplier's information do not nil")
	}

	if input.Email == "" {
		return errors.New("the supplier's information do not nil")
	}

	return nil
}

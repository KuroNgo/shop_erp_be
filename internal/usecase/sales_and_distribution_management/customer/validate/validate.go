package validate

import (
	"errors"
	customerdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/customer"
)

func Customer(input *customerdomain.Input) error {
	if input.FirstName == "" {
		return errors.New("the customer's information is invalid")
	}

	if input.LastName == "" {
		return errors.New("the customer's information is invalid")
	}

	if input.Email == "" {
		return errors.New("the customer's information is invalid")
	}

	if input.PhoneNumber == "" {
		return errors.New("the customer's information is invalid")
	}

	if input.Address == "" {
		return errors.New("the customer's information is invalid")
	}

	if input.City == "" {
		return errors.New("the customer's information is invalid")
	}

	if input.Country == "" {
		return errors.New("the customer's information is invalid")
	}
	return nil
}

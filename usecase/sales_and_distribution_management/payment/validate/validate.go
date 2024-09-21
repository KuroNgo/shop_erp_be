package validate

import (
	"errors"
	payments_domain "shop_erp_mono/domain/sales_and_distribution_management/payments"
)

func Payment(input *payments_domain.Input) error {
	if input.OrderID == "" {
		return errors.New("the payment's information is invalid")
	}

	if input.OrderID == "" {
		return errors.New("the payment's information is invalid")
	}

	if input.OrderID == "" {
		return errors.New("the payment's information is invalid")
	}

	if input.OrderID == "" {
		return errors.New("the payment's information is invalid")
	}

	if input.OrderID == "" {
		return errors.New("the payment's information is invalid")
	}
	return nil
}

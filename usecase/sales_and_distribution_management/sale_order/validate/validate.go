package validate

import (
	"errors"
	sale_orders_domain "shop_erp_mono/domain/sales_and_distribution_management/sale_orders"
)

func SaleOrder(input *sale_orders_domain.Input) error {
	if input.Status == "" {
		return errors.New("the payment's information is invalid")
	}

	if input.CustomerID == "" {
		return errors.New("the payment's information is invalid")
	}

	if input.ShippingAddress == "" {
		return errors.New("the payment's information is invalid")
	}

	if input.OrderNumber == "" {
		return errors.New("the payment's information is invalid")
	}

	if input.OrderDate.IsZero() {
		return errors.New("the payment's information is invalid")
	}

	return nil
}

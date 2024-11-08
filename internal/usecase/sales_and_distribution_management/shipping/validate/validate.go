package validate

import (
	"errors"
	shippingdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/shipping"
)

func Shipping(input *shippingdomain.Input) error {
	if input.OrderID == "" {
		return errors.New("the shipping's information is invalid")
	}

	if input.ShippingMethod == "" {
		return errors.New("the shipping's information is invalid")
	}

	if input.ShippingDate.IsZero() {
		return errors.New("the payment's information is invalid")
	}

	if input.EstimatedDelivery.IsZero() {
		return errors.New("the payment's information is invalid")
	}

	if input.ActualDelivery.IsZero() {
		return errors.New("the payment's information is invalid")
	}

	if input.TrackingNumber == "" {
		return errors.New("the payment's information is invalid")
	}

	if input.Status == "" {
		return errors.New("the payment's information is invalid")
	}

	return nil
}

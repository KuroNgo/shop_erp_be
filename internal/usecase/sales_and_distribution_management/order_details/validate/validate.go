package validate

import (
	"errors"
	order_details_domain "shop_erp_mono/internal/domain/sales_and_distribution_management/order_details"
)

func OrderDetail(input *order_details_domain.Input) error {
	if input.OrderID == "" {
		return errors.New("the order detail's information is invalid")
	}

	if input.ProductID == "" {
		return errors.New("the order detail's information is invalid")
	}

	if input.Quantity < 0 {
		return errors.New("the order detail's information is invalid")
	}

	if input.UnitPrice < 0 {
		return errors.New("the order detail's information is invalid")
	}

	return nil
}

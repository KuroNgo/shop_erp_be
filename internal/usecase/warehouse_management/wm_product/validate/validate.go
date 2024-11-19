package validate

import (
	"errors"
	productdomain "shop_erp_mono/internal/domain/warehouse_management/product"
)

func Product(input *productdomain.Input) error {
	if input.ProductName == "" {
		return errors.New("the wm_product's information do not nil")
	}

	if input.Category == "" {
		return errors.New("the wm_product's information do not nil")
	}
	if input.Description == "" {
		return errors.New("the wm_product's information do not nil")
	}

	if input.Price < 0 {
		return errors.New("the wm_product's information do not valid")
	}

	return nil
}

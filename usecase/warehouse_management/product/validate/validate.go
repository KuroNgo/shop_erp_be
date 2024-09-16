package validate

import (
	"errors"
	productdomain "shop_erp_mono/domain/warehouse_management/product"
)

func ValidateProduct(input *productdomain.Input) error {
	if input.ProductName == "" {
		return errors.New("the product's information do not nil")
	}

	if input.Category == "" {
		return errors.New("the product's information do not nil")
	}
	if input.Description == "" {
		return errors.New("the product's information do not nil")
	}

	if input.QuantityInStock < 0 {
		return errors.New("the product's information do not valid")
	}

	if input.Price < 0 {
		return errors.New("the product's information do not valid")
	}

	return nil
}

package validate

import (
	"errors"
	categorydomain "shop_erp_mono/domain/warehouse_management/product_category"
)

func Category(input *categorydomain.Input) error {
	if input.CategoryName == "" {
		return errors.New("the category's information do not nil")
	}

	if input.Description == "" {
		return errors.New("the category's information do not nil")
	}

	return nil
}

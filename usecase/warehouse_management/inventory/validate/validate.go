package validate

import (
	"errors"
	inventorydomain "shop_erp_mono/domain/warehouse_management/inventory"
)

func Inventory(input *inventorydomain.Input) error {
	if input.ProductName == "" {
		return errors.New("the inventory's information do not nil")
	}

	if input.WarehouseName == "" {
		return errors.New("the inventory's information do not nil")
	}

	if input.Quantity < 0 {
		return errors.New("the inventory's information is invalid")
	}

	return nil
}

package validate

import (
	"errors"
	warehousedomain "shop_erp_mono/internal/domain/warehouse_management/warehouse"
)

func Warehouse(input *warehousedomain.Input) error {
	if input.WarehouseName == "" {
		return errors.New("the warehouse's information is nil")
	}

	if input.Location == "" {
		return errors.New("the warehouse's information is nil")
	}

	if input.Capacity < 0 {
		return errors.New("the warehouse's information is invalid")
	}

	return nil
}

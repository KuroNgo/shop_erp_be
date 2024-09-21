package validate

import (
	"errors"
	stockmovementdomain "shop_erp_mono/domain/warehouse_management/stockmovement"
)

func StockMovement(input *stockmovementdomain.Input) error {
	if input.Product == "" {
		return errors.New("the stock movement's information do not nil")
	}

	if input.Warehouse == "" {
		return errors.New("the stock movement's information do not nil")
	}

	if input.Quantity < 0 {
		return errors.New("the stock movement's information do not valid")
	}

	if input.MovementType == "" {
		return errors.New("the stock movement's information do not nil")
	}

	if input.MovementDate.IsZero() {
		return errors.New("the stock movement's information do not nil")
	}

	if input.Reference == "" {
		return errors.New("the stock movement's information do not nil")
	}

	if input.User == "" {
		return errors.New("the stock movement's information do not nil")
	}

	return nil
}

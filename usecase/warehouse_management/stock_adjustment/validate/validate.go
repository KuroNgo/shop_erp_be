package validate

import (
	"errors"
	stockadjustmentdomain "shop_erp_mono/domain/warehouse_management/stock_adjustment"
)

func ValidateStockAdjustment(input *stockadjustmentdomain.Input) error {
	if input.Product == "" {
		return errors.New("the stock adjustment's information do not nil")
	}

	if input.Warehouse == "" {
		return errors.New("the stock adjustment's information do not nil")
	}

	if input.Quantity < 0 {
		return errors.New("the stock adjustment's information do not valid")
	}

	if input.Reason == "" {
		return errors.New("the stock adjustment's information do not nil")
	}

	if input.AdjustmentType == "" {
		return errors.New("the stock adjustment's information do not nil")
	}

	if input.AdjustmentDate.IsZero() {
		return errors.New("the stock adjustment's information do not nil")
	}

	return nil
}

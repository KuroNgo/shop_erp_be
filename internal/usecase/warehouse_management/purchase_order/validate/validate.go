package validate

import (
	"errors"
	purchaseorderdomain "shop_erp_mono/internal/domain/warehouse_management/purchase_order"
)

func PurchaseOrder(input *purchaseorderdomain.Input) error {
	if input.Supplier == "" {
		return errors.New("the purchase order's information do not nil")
	}

	if input.OrderNumber == "" {
		return errors.New("the purchase order's information do not nil")
	}

	if input.OrderDate.IsZero() {
		return errors.New("the purchase order's information do not nil")
	}

	if input.TotalAmount < 0 {
		return errors.New("the purchase order's information do not nil")
	}

	if input.Status == "" {
		return errors.New("the purchase order's information do not nil")
	}

	return nil
}

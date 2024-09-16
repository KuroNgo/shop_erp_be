package validate

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	purchaseorderdetaildomain "shop_erp_mono/domain/warehouse_management/purchase_order_detail"
)

func ValidatePurchaseOrderDetail(input *purchaseorderdetaildomain.Input) error {
	if input.Product == "" {
		return errors.New("the purchase order detail's information do not nil")
	}

	if input.PurchaseOrderID == primitive.NilObjectID {
		return errors.New("the purchase order detail's information do not nil")
	}

	if input.Quantity < 0 {
		return errors.New("the purchase order detail's information do not nil")
	}

	if input.UnitPrice < 0 {
		return errors.New("the purchase order detail's information do not nil")
	}

	return nil
}

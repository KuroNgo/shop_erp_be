package purchase_order_controller

import (
	"shop_erp_mono/bootstrap"
	purchaseorderdomain "shop_erp_mono/domain/warehouse_management/purchase_order"
)

type PurchaseOrderController struct {
	Database             *bootstrap.Database
	PurchaseOrderUseCase purchaseorderdomain.IPurchaseOrderUseCase
}

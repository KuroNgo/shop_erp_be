package purchase_order_controller

import (
	"shop_erp_mono/internal/config"
	purchaseorderdomain "shop_erp_mono/internal/domain/warehouse_management/purchase_order"
)

type PurchaseOrderController struct {
	Database             *config.Database
	PurchaseOrderUseCase purchaseorderdomain.IPurchaseOrderUseCase
}

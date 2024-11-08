package purchase_order_detail_controller

import (
	"shop_erp_mono/internal/config"
	purchaseorderdetaildomain "shop_erp_mono/internal/domain/warehouse_management/purchase_order_detail"
)

type PurchaseOrderDetailController struct {
	Database                   *config.Database
	PurchaseOrderDetailUseCase purchaseorderdetaildomain.IPurchaseOrderDetailUseCase
}

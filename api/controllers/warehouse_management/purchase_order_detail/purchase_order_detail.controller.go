package purchase_order_detail_controller

import (
	"shop_erp_mono/bootstrap"
	purchaseorderdetaildomain "shop_erp_mono/domain/warehouse_management/purchase_order_detail"
)

type PurchaseOrderDetailController struct {
	Database                   *bootstrap.Database
	PurchaseOrderDetailUseCase purchaseorderdetaildomain.IPurchaseOrderDetailUseCase
}

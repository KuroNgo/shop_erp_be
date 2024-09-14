package inventory_controller

import (
	"shop_erp_mono/bootstrap"
	inventorydomain "shop_erp_mono/domain/warehouse_management/inventory"
)

type InventoryController struct {
	Database         *bootstrap.Database
	InventoryUseCase inventorydomain.InventoryUseCase
}

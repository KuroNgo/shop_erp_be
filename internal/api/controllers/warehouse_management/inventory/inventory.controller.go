package inventory_controller

import (
	"shop_erp_mono/internal/config"
	inventorydomain "shop_erp_mono/internal/domain/warehouse_management/inventory"
)

type InventoryController struct {
	Database         *config.Database
	InventoryUseCase inventorydomain.InventoryUseCase
}

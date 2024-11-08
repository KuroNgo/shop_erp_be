package warehouse_controller

import (
	"shop_erp_mono/internal/config"
	warehouse_domain "shop_erp_mono/internal/domain/warehouse_management/warehouse"
)

type WarehouseController struct {
	Database         *config.Database
	WarehouseUseCase warehouse_domain.IWarehouseUseCase
}

package warehouse_controller

import (
	"shop_erp_mono/bootstrap"
	warehouse_domain "shop_erp_mono/domain/warehouse_management/warehouse"
)

type WarehouseController struct {
	Database         *bootstrap.Database
	WarehouseUseCase warehouse_domain.IWarehouseUseCase
}

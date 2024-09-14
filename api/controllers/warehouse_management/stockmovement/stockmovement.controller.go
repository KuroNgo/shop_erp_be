package stockmovement_controller

import (
	"shop_erp_mono/bootstrap"
	stockmovement_domain "shop_erp_mono/domain/warehouse_management/stockmovement"
)

type StockMovementController struct {
	Database             *bootstrap.Database
	StockMovementUseCase stockmovement_domain.IStockMovementUseCase
}

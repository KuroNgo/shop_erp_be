package stockmovement_controller

import (
	"shop_erp_mono/internal/config"
	stockmovement_domain "shop_erp_mono/internal/domain/warehouse_management/stockmovement"
)

type StockMovementController struct {
	Database             *config.Database
	StockMovementUseCase stockmovement_domain.IStockMovementUseCase
}

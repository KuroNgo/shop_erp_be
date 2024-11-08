package stock_adjustment_controller

import (
	"shop_erp_mono/internal/config"
	"shop_erp_mono/internal/domain/warehouse_management/stock_adjustment"
)

type StockAdjustmentController struct {
	Database               *config.Database
	StockAdjustmentUseCase stock_adjustment_domain.IStockAdjustmentUseCase
}

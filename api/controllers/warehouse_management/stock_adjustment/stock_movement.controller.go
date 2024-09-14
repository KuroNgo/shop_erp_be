package stock_adjustment_controller

import (
	"shop_erp_mono/bootstrap"
	"shop_erp_mono/domain/warehouse_management/stock_adjustment"
)

type StockAdjustmentController struct {
	Database               *bootstrap.Database
	StockAdjustmentUseCase stock_adjustment_domain.IStockAdjustmentUseCase
}

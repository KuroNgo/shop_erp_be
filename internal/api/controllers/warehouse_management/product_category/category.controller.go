package category_controller

import (
	"shop_erp_mono/internal/config"
	categorydomain "shop_erp_mono/internal/domain/warehouse_management/product_category"
)

type CategoryController struct {
	Database        *config.Database
	CategoryUseCase categorydomain.ICategoryUseCase
}

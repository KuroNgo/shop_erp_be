package category_controller

import (
	"shop_erp_mono/bootstrap"
	categorydomain "shop_erp_mono/domain/warehouse_management/product_category"
)

type CategoryController struct {
	Database        *bootstrap.Database
	CategoryUseCase categorydomain.ICategoryUseCase
}

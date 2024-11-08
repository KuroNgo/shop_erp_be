package product_controller

import (
	"shop_erp_mono/internal/config"
	productdomain "shop_erp_mono/internal/domain/warehouse_management/product"
)

type ProductController struct {
	Database       *config.Database
	ProductUseCase productdomain.IProductUseCase
}

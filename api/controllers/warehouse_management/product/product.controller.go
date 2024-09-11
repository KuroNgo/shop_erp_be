package product_controller

import (
	"shop_erp_mono/bootstrap"
	productdomain "shop_erp_mono/domain/warehouse_management/product"
)

type ProductController struct {
	Database       *bootstrap.Database
	ProductUseCase productdomain.IProductUseCase
}

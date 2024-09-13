package supplier_controller

import (
	"shop_erp_mono/bootstrap"
	supplierdomain "shop_erp_mono/domain/warehouse_management/supplier"
)

type SupplierController struct {
	Database        *bootstrap.Database
	SupplierUseCase supplierdomain.ISupplierUseCase
}

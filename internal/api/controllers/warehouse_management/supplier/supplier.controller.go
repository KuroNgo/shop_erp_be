package supplier_controller

import (
	"shop_erp_mono/internal/config"
	supplierdomain "shop_erp_mono/internal/domain/warehouse_management/supplier"
)

type SupplierController struct {
	Database        *config.Database
	SupplierUseCase supplierdomain.ISupplierUseCase
}

package contract_controller

import (
	"shop_erp_mono/bootstrap"
	contractsdomain "shop_erp_mono/domain/human_resource_management/contracts"
)

type ContractController struct {
	Database        *bootstrap.Database
	ContractUseCase contractsdomain.IContractsUseCase
}

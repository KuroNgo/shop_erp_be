package contract_controller

import (
	"shop_erp_mono/internal/config"
	contractsdomain "shop_erp_mono/internal/domain/human_resource_management/contracts"
)

type ContractController struct {
	Database        *config.Database
	ContractUseCase contractsdomain.IContractsUseCase
}

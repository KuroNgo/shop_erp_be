package benefit_controller

import (
	"shop_erp_mono/internal/config"
	benefitsdomain "shop_erp_mono/internal/domain/human_resource_management/benefits"
)

type BenefitController struct {
	Database       *config.Database
	BenefitUseCase benefitsdomain.IBenefitUseCase
}

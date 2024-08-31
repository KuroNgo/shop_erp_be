package benefit_controller

import (
	"shop_erp_mono/bootstrap"
	benefitsdomain "shop_erp_mono/domain/human_resource_management/benefits"
)

type BenefitController struct {
	Database       *bootstrap.Database
	BenefitUseCase benefitsdomain.IBenefitUseCase
}

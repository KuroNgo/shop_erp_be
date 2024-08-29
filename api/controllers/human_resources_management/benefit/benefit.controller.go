package benefit_controller

import (
	"shop_erp_mono/bootstrap"
	benefits_domain "shop_erp_mono/domain/human_resource_management/benefits"
)

type BenefitController struct {
	Database       *bootstrap.Database
	BenefitUseCase benefits_domain.IBenefitUseCase
}

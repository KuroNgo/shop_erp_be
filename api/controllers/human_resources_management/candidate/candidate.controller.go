package candidate_controller

import (
	"shop_erp_mono/bootstrap"
	candidatedomain "shop_erp_mono/domain/human_resource_management/candidate"
)

type CandidateController struct {
	Database         *bootstrap.Database
	CandidateUseCase candidatedomain.ICandidateUseCase
}

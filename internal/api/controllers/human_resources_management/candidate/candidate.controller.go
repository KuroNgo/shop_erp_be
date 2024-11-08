package candidate_controller

import (
	"shop_erp_mono/internal/config"
	candidatedomain "shop_erp_mono/internal/domain/human_resource_management/candidate"
)

type CandidateController struct {
	Database         *config.Database
	CandidateUseCase candidatedomain.ICandidateUseCase
}

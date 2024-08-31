package performance_review_controller

import (
	"shop_erp_mono/bootstrap"
	performance_review_domain "shop_erp_mono/domain/human_resource_management/performance_review"
)

type PerformanceReviewController struct {
	Database                 *bootstrap.Database
	PerformanceReviewUseCase performance_review_domain.IPerformanceReviewUseCase
}

package performance_review_controller

import (
	"shop_erp_mono/internal/config"
	performance_review_domain "shop_erp_mono/internal/domain/human_resource_management/performance_review"
)

type PerformanceReviewController struct {
	Database                 *config.Database
	PerformanceReviewUseCase performance_review_domain.IPerformanceReviewUseCase
}

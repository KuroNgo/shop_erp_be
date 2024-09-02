package performance_review_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	performancereviewdomain "shop_erp_mono/domain/human_resource_management/performance_review"
)

// CreateOnePerformanceReview create the performance review's information
// @Summary Create Performance Review Information
// @Description Create the performance review's information
// @Tags Performance Review
// @Accept json
// @Produce json
// @Param PerformanceReview body performance_review_domain.Input true "Performance Review data"
// @Router /api/v1/performance_reviews/create [post]
// @Security CookieAuth
func (p *PerformanceReviewController) CreateOnePerformanceReview(ctx *gin.Context) {
	var input performancereviewdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	err := p.PerformanceReviewUseCase.CreateOne(ctx, &input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

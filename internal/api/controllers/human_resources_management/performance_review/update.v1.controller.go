package performance_review_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	performancereviewdomain "shop_erp_mono/internal/domain/human_resource_management/performance_review"
)

// UpdateOneWithEmailEmployee update the performance review's information
// @Summary Update Performance Review Information
// @Description Update the performance review's information
// @Tags Performance Review
// @Accept json
// @Produce json
// @Param _id query string true "Performance Review  ID"
// @Param PerformanceReview body performance_review_domain.Input1 true "Performance Review data"
// @Router /api/v1/performance-reviews/update [put]
// @Security CookieAuth
func (p *PerformanceReviewController) UpdateOneWithEmailEmployee(ctx *gin.Context) {
	var input performancereviewdomain.Input1
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	_id := ctx.Query("_id")

	err := p.PerformanceReviewUseCase.UpdateOneWithEmailEmployee(ctx, _id, &input)
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

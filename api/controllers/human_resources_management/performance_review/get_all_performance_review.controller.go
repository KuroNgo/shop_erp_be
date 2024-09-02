package performance_review_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllPerformanceReview Get all the performance review's information
// @Summary Get all Performance Review Information
// @Description Get all the performance review's information
// @Tags Performance Review
// @Accept json
// @Produce json
// @Router /api/v1/performance_reviews/get/all [get]
// @Security CookieAuth
func (p *PerformanceReviewController) GetAllPerformanceReview(ctx *gin.Context) {
	data, err := p.PerformanceReviewUseCase.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}

package performance_review_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetOneByEmailPerformanceReview Get the performance review's information
// @Summary Get Performance Review Information
// @Description Get the performance review's information
// @Tags Performance Review
// @Accept json
// @Produce json
// @Param email path string true "Performance Review  ID"
// @Router /api/v1/performance_reviews/get/email [get]
// @Security CookieAuth
func (p *PerformanceReviewController) GetOneByEmailPerformanceReview(ctx *gin.Context) {
	email := ctx.Param("email")

	data, err := p.PerformanceReviewUseCase.GetOneByEmailEmployee(ctx, email)
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

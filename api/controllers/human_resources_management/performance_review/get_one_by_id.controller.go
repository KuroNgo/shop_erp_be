package performance_review_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetOneByIDPerformanceReview Get the performance review's information
// @Summary Get Performance Review Information
// @Description Get the performance review's information
// @Tags Performance Review
// @Accept json
// @Produce json
// @Param _id path string true "Performance Review  ID"
// @Router /api/v1/performance_reviews/get/_id [get]
// @Security CookieAuth
func (p *PerformanceReviewController) GetOneByIDPerformanceReview(ctx *gin.Context) {
	id := ctx.Query("id")

	data, err := p.PerformanceReviewUseCase.GetOneByID(ctx, id)
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

package performance_review_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByEmailEmployee Get the performance review's information
// @Summary Get Performance Review Information
// @Description Get the performance review's information
// @Tags Performance Review
// @Accept json
// @Produce json
// @Param email query string true "Performance Review  ID"
// @Router /api/v1/performance-reviews/get/email [get]
// @Security CookieAuth
func (p *PerformanceReviewController) GetByEmailEmployee(ctx *gin.Context) {
	email := ctx.Query("email")

	data, err := p.PerformanceReviewUseCase.GetByEmailEmployee(ctx, email)
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

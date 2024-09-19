package performance_review_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOne delete the performance review's information
// @Summary Delete Performance Review Information
// @Description Delete the performance review's information
// @Tags Performance Review
// @Accept json
// @Produce json
// @Param _id path string true "Performance Review  ID"
// @Router /api/v1/performance_reviews/delete [delete]
// @Security CookieAuth
func (p *PerformanceReviewController) DeleteOne(ctx *gin.Context) {
	id := ctx.Query("_id")

	err := p.PerformanceReviewUseCase.DeleteOne(ctx, id)
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

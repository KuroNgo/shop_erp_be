package performance_review_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (p *PerformanceReviewController) DeleteOnePerformanceReview(ctx *gin.Context) {
	id := ctx.Param("_id")

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

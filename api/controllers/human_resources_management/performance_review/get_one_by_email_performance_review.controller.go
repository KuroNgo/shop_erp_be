package performance_review_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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

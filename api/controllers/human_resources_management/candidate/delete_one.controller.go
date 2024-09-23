package candidate_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *CandidateController) DeleteOne(ctx *gin.Context) {
	candidateID := ctx.Query("_id")

	if err := c.CandidateUseCase.DeleteOne(ctx, candidateID); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

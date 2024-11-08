package candidate_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	candidatedomain "shop_erp_mono/internal/domain/human_resource_management/candidate"
)

func (e *CandidateController) UpdateOne(ctx *gin.Context) {
	var input candidatedomain.Candidate
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	candidateID := ctx.Query("_id")

	err := e.CandidateUseCase.UpdateOne(ctx, candidateID, &input)
	if err != nil {
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

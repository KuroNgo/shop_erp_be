package candidate_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Status struct {
	Status string `bson:"status" json:"status"`
}

func (e *CandidateController) UpdateStatus(ctx *gin.Context) {
	candidateID := ctx.Query("_id")
	var status Status
	if err := ctx.ShouldBindJSON(&status); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	err := e.CandidateUseCase.UpdateStatus(ctx, candidateID, status.Status)
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

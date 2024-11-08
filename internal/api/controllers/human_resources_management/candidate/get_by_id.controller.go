package candidate_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (e *CandidateController) GetByID(ctx *gin.Context) {
	employeeID := ctx.Query("_id")

	data, err := e.CandidateUseCase.GetByID(ctx, employeeID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
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

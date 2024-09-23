package candidate_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (e *CandidateController) GetByName(ctx *gin.Context) {
	name := ctx.Query("name")

	data, err := e.CandidateUseCase.GetByID(ctx, name)
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

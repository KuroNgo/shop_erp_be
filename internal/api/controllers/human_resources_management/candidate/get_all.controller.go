package candidate_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *CandidateController) GetAll(ctx *gin.Context) {
	data, err := c.CandidateUseCase.GetAll(ctx)
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

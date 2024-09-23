package candidate_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	candidatedomain "shop_erp_mono/domain/human_resource_management/candidate"
)

func (c *CandidateController) CreateOne(ctx *gin.Context) {
	var input candidatedomain.Candidate
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	err := c.CandidateUseCase.CreateOne(ctx, &input)
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

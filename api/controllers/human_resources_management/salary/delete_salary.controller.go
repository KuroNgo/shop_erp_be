package salary_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *SalaryController) DeleteOneRole(ctx *gin.Context) {
	id := ctx.Param("_id")

	err := s.SalaryUseCase.DeleteOne(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

package salary_base

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *BaseSalaryController) GetByID(ctx *gin.Context) {
	id := ctx.Query("_id")

	data, err := s.BaseSalaryUseCase.GetByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
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

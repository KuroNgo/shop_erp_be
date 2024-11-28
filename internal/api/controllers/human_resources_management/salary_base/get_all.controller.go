package salary_base

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *BaseSalaryController) GetAll(ctx *gin.Context) {
	data, err := s.BaseSalaryUseCase.GetAll(ctx)
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

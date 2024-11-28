package salary_base

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAll retrieves the base salary's information
// @Summary Get Base Salary Information
// @Description Retrieves the base salary's information
// @Tags Base Salary
// @Accept  json
// @Produce  json
// @Router /api/v1/base-salaries/get/all [get]
// @Security CookieAuth
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

package salary_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAll retrieves the salary's information
// @Summary Get Salary Information
// @Description Retrieves the salary's information
// @Tags Salary
// @Accept  json
// @Produce  json
// @Router /api/v1/salaries/get/all [get]
// @Security CookieAuth
func (s *SalaryController) GetAll(ctx *gin.Context) {
	data, err := s.SalaryUseCase.GetAll(ctx)
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

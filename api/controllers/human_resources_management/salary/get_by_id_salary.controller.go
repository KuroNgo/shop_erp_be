package salary_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetOneSalaryByID retrieves the salary's information
// @Summary Get Salary Information By ID
// @Description Retrieves the salary's information id
// @Tags Salary
// @Accept  json
// @Produce  json
// @Router /api/v1/salaries/get/one/_id [get]
// @Security CookieAuth
func (s *SalaryController) GetOneSalaryByID(ctx *gin.Context) {
	id := ctx.Param("_id")

	data, err := s.SalaryUseCase.GetOneByID(ctx, id)
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
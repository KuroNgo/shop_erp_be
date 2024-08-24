package salary_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllSalary retrieves the salary's information
// @Summary Get Salary Information
// @Description Retrieves the salary's information
// @Tags Salary
// @Accept  json
// @Produce  json
// @Success 200 {object} salary_domain.Salary
// @Failure 400 {object} map[string]interface{} "status: fail, message: detailed error message"
// @Failure 401 {object} map[string]interface{} "status: fail, message: You are not logged in!"
// @Router /api/v1/all/salaries [get]
// @Security CookieAuth
func (s *SalaryController) GetAllSalary(ctx *gin.Context) {
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

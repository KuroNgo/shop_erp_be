package salary_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetOneSalaryByRole retrieves the salary's information
// @Summary Get Salary Information By Role
// @Description Retrieves the salary's information role
// @Tags Salary
// @Accept  json
// @Produce  json
// @Router /api/v1/salaries/get/one/role [get]
// @Security CookieAuth
func (s *SalaryController) GetOneSalaryByRole(ctx *gin.Context) {
	role := ctx.Param("role")

	data, err := s.SalaryUseCase.GetOneByRole(ctx, role)
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

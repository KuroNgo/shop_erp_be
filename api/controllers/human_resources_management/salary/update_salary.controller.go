package salary_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	salarydomain "shop_erp_mono/domain/human_resource_management/salary"
)

// UpdateOneSalary updates the salary's information
// @Summary Update Salary Information
// @Description Updates the salary's information
// @Tags Salary
// @Accept json
// @Produce json
// @Router /api/v1/salaries/update [patch]
// @Security CookieAuth
func (s *SalaryController) UpdateOneSalary(ctx *gin.Context) {
	var salary salarydomain.Input
	if err := ctx.ShouldBindJSON(&salary); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if err := s.SalaryUseCase.UpdateOne(ctx, &salary); err != nil {
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

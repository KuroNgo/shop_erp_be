package salary_base

import (
	"github.com/gin-gonic/gin"
	"net/http"
	basesalarydomain "shop_erp_mono/internal/domain/human_resource_management/salary_base"
)

// UpdateOne updates the base salary's information
// @Summary Update Salary Information
// @Description Updates the salary's information
// @Tags Base Salary
// @Accept json
// @Produce json
// @Param Salary body base_salary_domain.Input true "Salary data"
// @Param _id path string true "Employee ID"
// @Router /api/v1/salaries/update [put]
// @Security CookieAuth
func (s *BaseSalaryController) UpdateOne(ctx *gin.Context) {
	var salary basesalarydomain.Input
	if err := ctx.ShouldBindJSON(&salary); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	_id := ctx.Query("_id")

	if err := s.BaseSalaryUseCase.UpdateOne(ctx, _id, &salary); err != nil {
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

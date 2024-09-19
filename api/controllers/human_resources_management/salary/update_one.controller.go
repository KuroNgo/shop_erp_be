package salary_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	salarydomain "shop_erp_mono/domain/human_resource_management/salary"
)

// UpdateOne updates the salary's information
// @Summary Update Salary Information
// @Description Updates the salary's information
// @Tags Salary
// @Accept json
// @Produce json
// @Param Salary body salary_domain.Salary true "Salary data"
// @Param _id path string true "Employee ID"
// @Router /api/v1/salaries/update [put]
// @Security CookieAuth
func (s *SalaryController) UpdateOne(ctx *gin.Context) {
	var salary salarydomain.Input
	if err := ctx.ShouldBindJSON(&salary); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	_id := ctx.Query("_id")

	if err := s.SalaryUseCase.UpdateOne(ctx, _id, &salary); err != nil {
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

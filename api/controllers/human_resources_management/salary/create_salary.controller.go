package salary_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	salarydomain "shop_erp_mono/domain/human_resource_management/salary"
)

// CreateOneSalary Create a new salary
// @Summary Create salary
// @Description Create new salary
// @Tags Salary
// @Accept json
// @Produce json
// @Param Role body salary_domain.Salary true "Salary data"
// @Security ApiKeyAuth
// @Router /api/v1/salaries/create [post]
func (s *SalaryController) CreateOneSalary(ctx *gin.Context) {
	var salary salarydomain.Input
	if err := ctx.ShouldBindJSON(&salary); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if err := s.SalaryUseCase.CreateOne(ctx, &salary); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

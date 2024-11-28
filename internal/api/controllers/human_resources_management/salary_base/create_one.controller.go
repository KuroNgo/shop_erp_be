package salary_base

import (
	"github.com/gin-gonic/gin"
	"net/http"
	basesalarydomain "shop_erp_mono/internal/domain/human_resource_management/salary_base"
)

func (s *BaseSalaryController) CreateOne(ctx *gin.Context) {
	var salary basesalarydomain.Input
	if err := ctx.ShouldBindJSON(&salary); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if err := s.BaseSalaryUseCase.CreateOne(ctx, &salary); err != nil {
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

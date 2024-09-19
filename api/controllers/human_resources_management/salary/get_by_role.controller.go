package salary_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByRoleTitle retrieves the salary's information
// @Summary Get Salary Information By Role
// @Description Retrieves the salary's information role
// @Tags Salary
// @Accept  json
// @Produce  json
// @Param role path string true "Role"
// @Router /api/v1/salaries/get/role [get]
// @Security CookieAuth
func (s *SalaryController) GetByRoleTitle(ctx *gin.Context) {
	role := ctx.Query("role")

	data, err := s.SalaryUseCase.GetByRoleTitle(ctx, role)
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

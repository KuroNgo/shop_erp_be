package salary_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOneSalary delete the salary's information
// @Summary Delete Salary Information
// @Description Deletes the salary's information
// @Tags Salary
// @Accept json
// @Produce json
// @Param _id path string true "Role ID"
// @Router /api/v1/salaries/delete [delete]
// @Security CookieAuth
func (s *SalaryController) DeleteOneSalary(ctx *gin.Context) {
	id := ctx.Query("_id")

	err := s.SalaryUseCase.DeleteOne(ctx, id)
	if err != nil {
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

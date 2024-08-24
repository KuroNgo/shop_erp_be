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
// @@Success 200 {object} "status: success, message: delete salary success"
// @Failure 400 {object} map[string]interface{} "status: fail, message: detailed error message"
// @Failure 401 {object} map[string]interface{} "status: fail, message: You are not logged in!"
// @Router /api/v1/salaries/delete [delete]
// @Security CookieAuth
func (s *SalaryController) DeleteOneSalary(ctx *gin.Context) {
	id := ctx.Param("_id")

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

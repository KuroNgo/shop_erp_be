package salary_base

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOne delete the base salary's information
// @Summary Delete Base Salary Information
// @Description Deletes the base salary's information
// @Tags Base Salary
// @Accept json
// @Produce json
// @Param _id path string true "Base Salary ID"
// @Router /api/v1/base-salaries/delete [delete]
// @Security CookieAuth
func (s *BaseSalaryController) DeleteOne(ctx *gin.Context) {
	id := ctx.Query("_id")

	err := s.BaseSalaryUseCase.DeleteOne(ctx, id)
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

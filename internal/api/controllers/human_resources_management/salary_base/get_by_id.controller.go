package salary_base

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByID retrieves the base salary's information
// @Summary Get Base Salary Information By ID
// @Description Retrieves the base salary's information id
// @Tags Base Salary
// @Accept  json
// @Produce  json
// @Param _id path string true "Base Salary ID"
// @Router /api/v1/base-salaries/get/_id [get]
// @Security CookieAuth
func (s *BaseSalaryController) GetByID(ctx *gin.Context) {
	id := ctx.Query("_id")

	data, err := s.BaseSalaryUseCase.GetByID(ctx, id)
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

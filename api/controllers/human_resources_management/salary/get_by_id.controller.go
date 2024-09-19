package salary_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetOneSalaryByID retrieves the salary's information
// @Summary Get Salary Information By ID
// @Description Retrieves the salary's information id
// @Tags Salary
// @Accept  json
// @Produce  json
// @Param _id path string true "Role ID"
// @Router /api/v1/salaries/get/_id [get]
// @Security CookieAuth
func (s *SalaryController) GetOneSalaryByID(ctx *gin.Context) {
	id := ctx.Query("_id")

	data, err := s.SalaryUseCase.GetOneByID(ctx, id)
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

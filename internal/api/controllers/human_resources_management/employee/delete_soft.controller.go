package employee_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteSoft delete the employee's information
// @Summary Delete Employee Information
// @Description Deletes the employee's information
// @Tags Employee
// @Produce json
// @Param _id query string true "Employee ID"
// @Router /api/v1/employees/delete [patch]
// @Security CookieAuth
func (e *EmployeeController) DeleteSoft(ctx *gin.Context) {
	attendanceID := ctx.Query("_id")

	if err := e.EmployeeUseCase.DeleteSoft(ctx, attendanceID); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

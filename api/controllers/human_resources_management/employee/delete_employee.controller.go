package employee_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOneEmployee delete the employee's information
// @Summary Delete Employee Information
// @Description Deletes the employee's information
// @Tags Employee
// @Produce json
// @Param _id path string true "Employee ID"
// @Router /api/v1/employees/delete [delete]
// @Security CookieAuth
func (e *EmployeeController) DeleteOneEmployee(ctx *gin.Context) {
	attendanceID := ctx.Param("_id")

	if err := e.EmployeeUseCase.DeleteOne(ctx, attendanceID); err != nil {
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
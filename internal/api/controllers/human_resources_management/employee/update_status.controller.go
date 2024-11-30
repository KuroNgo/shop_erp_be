package employee_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpdateStatus update the employee's information
// @Summary Update Employee Information
// @Description Create the employee's information
// @Tags Employee
// @Produce json
// @Param _id query string true "Employee ID"
// @Param status body employees_domain.Input true "Employee data"
// @Router /api/v1/employees/update/status [put]
// @Security CookieAuth
func (e *EmployeeController) UpdateStatus(ctx *gin.Context) {
	employeeID := ctx.Query("_id")
	status := ctx.Query("status")

	err := e.EmployeeUseCase.UpdateStatus(ctx, employeeID, status)
	if err != nil {
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

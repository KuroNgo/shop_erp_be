package employee_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByName retrieves the employee's information
// @Summary Get Employee Information
// @Description Retrieves the employee's information
// @Tags Employee
// @Produce  json
// @Param name query string true "Employee"
// @Router /api/v1/employees/get/name [get]
// @Security CookieAuth
func (e *EmployeeController) GetByName(ctx *gin.Context) {
	employeeEmail := ctx.Query("name")

	data, err := e.EmployeeUseCase.GetByName(ctx, employeeEmail)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
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

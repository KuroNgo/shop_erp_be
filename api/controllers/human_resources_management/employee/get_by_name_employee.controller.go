package employee_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// FetchByNameEmployee retrieves the employee's information
// @Summary Get Employee Information
// @Description Retrieves the employee's information
// @Tags Employee
// @Produce  json
// @Param name path string true "Employee"
// @Router /api/v1/employees/get/name [get]
// @Security CookieAuth
func (e *EmployeeController) FetchByNameEmployee(ctx *gin.Context) {
	employeeName := ctx.Param("name")

	data, err := e.EmployeeUseCase.GetOneByName(ctx, employeeName)
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

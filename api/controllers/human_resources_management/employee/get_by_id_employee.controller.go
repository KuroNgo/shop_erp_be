package employee_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// FetchByIDEmployee retrieves the employee's information
// @Summary Get Employee Information
// @Description Retrieves the employee's information
// @Tags Employee
// @Produce  json
// @Param _id path string true "Employee ID"
// @Router /api/v1/employees/get/_id [get]
// @Security CookieAuth
func (e *EmployeeController) FetchByIDEmployee(ctx *gin.Context) {
	employeeID := ctx.Param("_id")

	data, err := e.EmployeeUseCase.GetOneByID(ctx, employeeID)
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
package employee_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByID retrieves the employee's information
// @Summary Get Employee Information
// @Description Retrieves the employee's information
// @Tags Employee
// @Produce  json
// @Param _id query string true "Employee ID"
// @Router /api/v1/employees/get/_id [get]
// @Security CookieAuth
func (e *EmployeeController) GetByID(ctx *gin.Context) {
	employeeID := ctx.Query("_id")

	data, err := e.EmployeeUseCase.GetByID(ctx, employeeID)
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

package employee_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByEmail retrieves the employee's information
// @Summary Get Employee Information
// @Description Retrieves the employee's information
// @Tags Employee
// @Produce  json
// @Param email query string true "Employee Email"
// @Router /api/v1/employees/get/email [get]
// @Security CookieAuth
func (e *EmployeeController) GetByEmail(ctx *gin.Context) {
	email := ctx.Query("email")

	data, err := e.EmployeeUseCase.GetByEmail(ctx, email)
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

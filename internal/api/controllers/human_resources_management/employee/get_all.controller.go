package employee_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAll retrieves the employee's information
// @Summary Get Employee Information
// @Description Retrieves the employee's information
// @Tags Employee
// @Accept  json
// @Produce  json
// @Router /api/v1/employees/get/all [get]
// @Security CookieAuth
func (e *EmployeeController) GetAll(ctx *gin.Context) {
	data, err := e.EmployeeUseCase.GetAll(ctx)
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

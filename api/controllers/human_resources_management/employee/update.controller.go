package employee_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
)

// UpdateOne create the employee's information
// @Summary Create Employee Information
// @Description Create the employee's information
// @Tags Employee
// @Produce json
// @Param _id path string true "Employee ID"
// @Param attendance body employees_domain.Input true "Employee data"
// @Router /api/v1/employees/update [put]
// @Security CookieAuth
func (e *EmployeeController) UpdateOne(ctx *gin.Context) {
	var input employeesdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	employeeID := ctx.Query("_id")

	err := e.EmployeeUseCase.UpdateOne(ctx, employeeID, &input)
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

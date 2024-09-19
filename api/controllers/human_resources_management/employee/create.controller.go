package employee_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	employees_domain "shop_erp_mono/domain/human_resource_management/employees"
)

// CreateOne create the employee's information
// @Summary Create Employee Information
// @Description Create the employee's information
// @Tags Employee
// @Accept json
// @Produce json
// @Param attendance body employees_domain.Input true "Employee data"
// @Router /api/v1/employees/create [post]
// @Security CookieAuth
func (e *EmployeeController) CreateOne(ctx *gin.Context) {
	var input employees_domain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	err := e.EmployeeUseCase.CreateOne(ctx, &input)
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

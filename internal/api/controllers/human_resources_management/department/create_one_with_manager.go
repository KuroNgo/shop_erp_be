package department_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	departmentsdomain "shop_erp_mono/internal/domain/human_resource_management/departments"
	employeesdomain "shop_erp_mono/internal/domain/human_resource_management/employees"
)

type DepartmentWithManagerInput struct {
	Department departmentsdomain.Input `json:"department"`
	Manager    employeesdomain.Input   `json:"manager"`
}

// CreateOneWithManager create the department's information
// @Summary Create Department Information
// @Description Create the department's information along with a manager
// @Tags Department
// @Accept json
// @Produce json
// @Param Input body DepartmentWithManagerInput true "Department and Manager data"
// @Router /api/v1/departments/create/manager [post]
// @Security CookieAuth
func (d *DepartmentController) CreateOneWithManager(ctx *gin.Context) {
	var input DepartmentWithManagerInput
	// Bind JSON input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid input: " + err.Error(),
		})
		return
	}

	// Call use case to create department with manager
	err := d.DepartmentUseCase.CreateDepartmentWithManager(ctx, &input.Department, &input.Manager)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to create department: " + err.Error(),
		})
		return
	}

	// Return success response
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Department created successfully",
	})
}

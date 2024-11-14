package department_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpdateManager updates the department's information
// @Summary Update Department Information
// @Description Updates the department's information
// @Tags Department
// @Accept json
// @Produce json
// @Param department_id query string true "Department ID"
// @Param manager_id query string true "Manager ID"
// @Router /api/v1/departments/update [put]
// @Security CookieAuth
func (d *DepartmentController) UpdateManager(ctx *gin.Context) {
	departmentID := ctx.Query("department_id")
	managerID := ctx.Query("manager_id")

	if err := d.DepartmentUseCase.UpdateManager(ctx, departmentID, managerID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

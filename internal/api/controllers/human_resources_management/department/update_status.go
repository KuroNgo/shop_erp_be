package department_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpdateStatus updates the department's information
// @Summary Update Department Information
// @Description Updates the department's information
// @Tags Department
// @Accept json
// @Produce json
// @Router /api/v1/departments/update [put]
// @Security CookieAuth
func (d *DepartmentController) UpdateStatus(ctx *gin.Context) {
	currentUser, exist := ctx.Get("currentUser")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "You are not login!",
		})
		return
	}

	departmentID := ctx.Query("_id")
	status := ctx.Query("status")

	if err := d.DepartmentUseCase.UpdateStatus(ctx, departmentID, status, fmt.Sprintf("%s", currentUser)); err != nil {
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

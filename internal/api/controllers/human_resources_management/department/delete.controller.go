package department_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOne delete the department's information
// @Summary Delete Department Information
// @Description Deletes the department's information
// @Tags Department
// @Accept json
// @Produce json
// @Param _id query string true "Department ID"
// @Router /api/v1/departments/delete/_id [delete]
// @Security CookieAuth
func (d *DepartmentController) DeleteOne(ctx *gin.Context) {
	currentUser, exist := ctx.Get("currentUser")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "You are not login!",
		})
		return
	}
	departmentID := ctx.Query("_id")

	if err := d.DepartmentUseCase.DeleteOne(ctx, departmentID, fmt.Sprintf("%s", currentUser)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

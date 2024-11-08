package department_controller

import (
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
	departmentID := ctx.Query("_id")

	if err := d.DepartmentUseCase.DeleteOne(ctx, departmentID); err != nil {
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

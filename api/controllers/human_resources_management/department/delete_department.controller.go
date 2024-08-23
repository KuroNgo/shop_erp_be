package department_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOneDepartment delete the department's information
// @Summary Delete Department Information
// @Description Deletes the department's information
// @Tags Department
// @Accept json
// @Produce json
// @@Success 200 {object} "status: success, message:update department success"
// @Failure 400 {object} map[string]interface{} "status: fail, message: detailed error message"
// @Failure 401 {object} map[string]interface{} "status: fail, message: You are not logged in!"
// @Router /api/v1/departments/delete [delete]
// @Security CookieAuth
func (d *DepartmentController) DeleteOneDepartment(ctx *gin.Context) {
	departmentID := ctx.Param("_id")

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

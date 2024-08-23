package department_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// FetchAllDepartment retrieves the department's information
// @Summary Get Department Information
// @Description Retrieves the department's information
// @Tags Department
// @Accept  json
// @Produce  json
// @Success 200 {object} []departments_domain.Department
// @Failure 400 {object} map[string]interface{} "status: fail, message: detailed error message"
// @Failure 401 {object} map[string]interface{} "status: fail, message: You are not logged in!"
// @Router /api/all/departments [get]
// @Security CookieAuth
func (d *DepartmentController) FetchAllDepartment(ctx *gin.Context) {
	data, err := d.DepartmentUseCase.GetAll(ctx)
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

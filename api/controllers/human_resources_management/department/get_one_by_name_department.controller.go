package department_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// FetchOneDepartmentByName retrieves the department's information
// @Summary Get Department Information By Name
// @Description Retrieves the department's information name
// @Tags Department
// @Accept  json
// @Produce  json
// @Success 200 {object} departments_domain.Department
// @Failure 400 {object} map[string]interface{} "status: fail, message: detailed error message"
// @Failure 401 {object} map[string]interface{} "status: fail, message: You are not logged in!"
// @Router /api/v1/one/departments/name [get]
// @Security CookieAuth
func (d *DepartmentController) FetchOneDepartmentByName(ctx *gin.Context) {
	departmentName := ctx.Param("name")

	data, err := d.DepartmentUseCase.GetOneByName(ctx, departmentName)
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

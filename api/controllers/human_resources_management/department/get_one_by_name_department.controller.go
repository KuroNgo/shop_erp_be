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
// @Router /api/v1/departments/get/one/name [get]
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

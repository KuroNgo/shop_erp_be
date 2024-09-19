package department_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByName retrieves the department's information
// @Summary Get Department Information By Name
// @Description Retrieves the department's information name
// @Tags Department
// @Accept  json
// @Produce  json
// @Param name path string true "Contract ID"
// @Router /api/v1/departments/get/name [get]
// @Security CookieAuth
func (d *DepartmentController) GetByName(ctx *gin.Context) {
	departmentName := ctx.Query("name")

	data, err := d.DepartmentUseCase.GetByName(ctx, departmentName)
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

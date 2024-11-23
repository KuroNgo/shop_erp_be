package department_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByStatus retrieves the department's information
// @Summary Get Department Information By ID
// @Description Retrieves the department's information name
// @Tags Department
// @Accept  json
// @Produce  json
// @Param status query string true "Status"
// @Router /api/v1/departments/get/status [get]
// @Security CookieAuth
func (d *DepartmentController) GetByStatus(ctx *gin.Context) {
	departmentID := ctx.Query("status")

	data, err := d.DepartmentUseCase.GetByID(ctx, departmentID)
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

package department_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllSoftDelete retrieves the department's information in soft delete list
// @Summary Get Department Information
// @Description Retrieves the department's information
// @Tags Department
// @Accept  json
// @Produce  json
// @Router /api/v1/departments/get/all/deleted [get]
// @Security CookieAuth
func (d *DepartmentController) GetAllSoftDelete(ctx *gin.Context) {
	data, err := d.DepartmentUseCase.GetAllSoftDelete(ctx)
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

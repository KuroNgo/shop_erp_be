package department_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// FetchOneDepartmentByID retrieves the department's information
// @Summary Get Department Information By ID
// @Description Retrieves the department's information name
// @Tags Department
// @Accept  json
// @Produce  json
// @Router /api/v1/departments/get/one/_id [get]
// @Security CookieAuth
func (d *DepartmentController) FetchOneDepartmentByID(ctx *gin.Context) {
	departmentID := ctx.Param("_id")

	data, err := d.DepartmentUseCase.GetOneByID(ctx, departmentID)
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

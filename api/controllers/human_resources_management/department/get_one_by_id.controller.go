package department_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByID retrieves the department's information
// @Summary Get Department Information By ID
// @Description Retrieves the department's information name
// @Tags Department
// @Accept  json
// @Produce  json
// @Param _id query string true "Contract ID"
// @Router /api/v1/departments/get [get]
// @Security CookieAuth
func (d *DepartmentController) GetByID(ctx *gin.Context) {
	departmentID := ctx.Query("_id")

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

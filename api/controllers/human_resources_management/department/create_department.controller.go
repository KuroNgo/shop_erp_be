package department_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	departmentsdomain "shop_erp_mono/domain/human_resource_management/departments"
)

// CreateOneDepartment create the department's information
// @Summary Create Department Information
// @Description Create the department's information
// @Tags Department
// @Accept json
// @Produce json
// @Router /api/v1/departments/create [post]
// @Security CookieAuth
func (d *DepartmentController) CreateOneDepartment(ctx *gin.Context) {
	var input departmentsdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	err := d.DepartmentUseCase.CreateOne(ctx, &input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

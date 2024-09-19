package department_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	departmentsdomain "shop_erp_mono/domain/human_resource_management/departments"
)

// UpdateOne updates the department's information
// @Summary Update Department Information
// @Description Updates the department's information
// @Tags Department
// @Accept json
// @Produce json
// @Router /api/v1/departments/update [put]
// @Security CookieAuth
func (d *DepartmentController) UpdateOne(ctx *gin.Context) {
	var input departmentsdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	departmentID := ctx.Query("_id")

	if err := d.DepartmentUseCase.UpdateOne(ctx, departmentID, &input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

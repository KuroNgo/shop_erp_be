package department_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	departmentsdomain "shop_erp_mono/internal/domain/human_resource_management/departments"
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
	currentUser, exist := ctx.Get("currentUser")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "You are not login!",
		})
		return
	}

	var input departmentsdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	departmentID := ctx.Query("_id")

	if err := d.DepartmentUseCase.UpdateOne(ctx, departmentID, &input, fmt.Sprintf("%s", currentUser)); err != nil {
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

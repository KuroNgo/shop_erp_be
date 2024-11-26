package department_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	departmentsdomain "shop_erp_mono/internal/domain/human_resource_management/departments"
	"shop_erp_mono/pkg/shared/constant"
)

// CreateOne create the department's information
// @Summary Create Department Information
// @Description Create the department's information
// @Tags Department
// @Accept json
// @Produce json
// @Param Department body departments_domain.Input true "Department data"
// @Router /api/v1/departments/create [post]
// @Security CookieAuth
func (d *DepartmentController) CreateOne(ctx *gin.Context) {
	currentUser, exists := ctx.Get("currentUser")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "fail",
			"message": constant.MsgAPIUnauthorized,
		})
		return
	}

	var input departmentsdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": constant.MsgAPIBadRequest,
		})
		return
	}

	err := d.DepartmentUseCase.CreateOne(ctx, &input, fmt.Sprintf("%s", currentUser))
	if err != nil {
		switch err.Error() {
		case constant.MsgForbidden:
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":  "error",
				"message": constant.MsgAPIForbidden,
			})
		case constant.MsgConflict:
			ctx.JSON(http.StatusConflict, gin.H{
				"status":  "error",
				"message": constant.MsgAPIConflict,
			})
		default:
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": constant.MsgAPIBadRequest,
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": constant.MsgDataCreationSuccess,
	})
}

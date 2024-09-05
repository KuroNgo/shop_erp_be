package attendance_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	attendancedomain "shop_erp_mono/domain/human_resource_management/attendance"
)

// CreateOneAttendance create the attendance's information
// @Summary Create Attendance Information
// @Description Create the attendance's information
// @Tags Attendance
// @tag.name Attendance
// @tag.description Attendance represents the attendance information of an employee.
// @Accept json
// @Produce json
// @Param LoginUserRequestDto body attendance_domain.Input true "Attendance data"
// @Router /api/v1/attendances/create [post]
// @Security CookieAuth
func (a *AttendanceController) CreateOneAttendance(ctx *gin.Context) {
	var input attendancedomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	err := a.AttendanceUseCase.CreateOne(ctx, &input)
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

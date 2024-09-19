package attendance_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	attendancedomain "shop_erp_mono/domain/human_resource_management/attendance"
)

// CreateOne create the attendance's information
// @Summary Create Attendance Information
// @Description Create the attendance's information
// @Tags Attendance
// @Accept json
// @Produce json
// @Param CreateOneAttendance body attendance_domain.Input true "Attendance data"
// @Router /api/v1/attendances/create [post]
// @Security CookieAuth
func (a *AttendanceController) CreateOne(ctx *gin.Context) {
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

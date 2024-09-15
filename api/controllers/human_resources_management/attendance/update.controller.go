package attendance_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	attendancedomain "shop_erp_mono/domain/human_resource_management/attendance"
)

// UpdateOneAttendance updates the attendance's information
// @Summary Update Attendance Information
// @Description Updates the attendance's information
// @Tags Attendance
// @Produce json
// @Param _id path string true "Attendance ID"
// @Param attendance body attendance_domain.Input true "Attendance data"
// @Router /api/v1/attendances/_id [put]
// @Security CookieAuth
func (a *AttendanceController) UpdateOneAttendance(ctx *gin.Context) {
	var input attendancedomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	attendanceID := ctx.Param("_id")

	if err := a.AttendanceUseCase.UpdateOne(ctx, attendanceID, &input); err != nil {
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

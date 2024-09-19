package attendance_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOneAttendance delete the attendance's information
// @Summary Delete Attendance Information
// @Description Deletes the attendance's information by ID
// @Tags Attendance
// @Produce json
// @Param _id path string true "Attendance ID"
// @Router /api/v1/attendances/_id [delete]
// @Security CookieAuth
func (a *AttendanceController) DeleteOneAttendance(ctx *gin.Context) {
	attendanceID := ctx.Query("_id")

	if err := a.AttendanceUseCase.DeleteOne(ctx, attendanceID); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

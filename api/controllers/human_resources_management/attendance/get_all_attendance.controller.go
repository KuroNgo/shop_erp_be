package attendance_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// FetchAllAttendance retrieves the attendance's information
// @Summary Get Attendance Information
// @Description Retrieves the attendance's information
// @Tags Attendance
// @Accept  json
// @Produce  json
// @Router /api/v1/attendances/get/all [get]
// @Security CookieAuth
func (a *AttendanceController) FetchAllAttendance(ctx *gin.Context) {
	data, err := a.AttendanceUseCase.GetAll(ctx)
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
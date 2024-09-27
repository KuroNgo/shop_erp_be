package attendance_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByID retrieves the attendance's information
// @Summary Get Attendance Information By ID
// @Description Retrieves the attendance's information name
// @Tags Attendance
// @Accept  json
// @Produce  json
// @Param _id query string true "Attendance ID"
// @Router /api/v1/attendances/get/_id [get]
// @Security CookieAuth
func (a *AttendanceController) GetByID(ctx *gin.Context) {
	attendanceID := ctx.Query("_id")

	data, err := a.AttendanceUseCase.GetByID(ctx, attendanceID)
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

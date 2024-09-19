package attendance_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByEmail retrieves the attendance's information
// @Summary Get Attendance Information By ID
// @Description Retrieves the attendance's information name
// @Tags Attendance
// @Produce  json
// @Param email path string true "Attendance ID"
// @Router /api/v1/attendances/get/email [get]
// @Security CookieAuth
func (a *AttendanceController) GetByEmail(ctx *gin.Context) {
	email := ctx.Query("email")

	data, err := a.AttendanceUseCase.GetByEmail(ctx, email)
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

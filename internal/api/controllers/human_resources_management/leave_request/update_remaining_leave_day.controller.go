package leave_request_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpdateRemainingLeaveDays update the leave request's information
// @Summary Delete Leave Request Information
// @Description Delete the leave request's information
// @Tags Leave Request
// @Accept json
// @Produce json
// @Router /api/v1/leave-requests/update/remaining [put]
// @Security CookieAuth
func (l *LeaveRequestController) UpdateRemainingLeaveDays(ctx *gin.Context) {
	l.LeaveRequestUseCase.StartSchedulerUpdateRemainingLeaveDays(l.CronJob)

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

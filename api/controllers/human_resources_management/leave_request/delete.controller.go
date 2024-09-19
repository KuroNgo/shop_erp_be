package leave_request_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOneLeaveRequest delete the leave request's information
// @Summary Delete Leave Request Information
// @Description Delete the leave request's information
// @Tags Leave Request
// @Accept json
// @Produce json
// @Param _id path string true "Leave Request ID"
// @Router /api/v1/leave_requests/delete [delete]
// @Security CookieAuth
func (l *LeaveRequestController) DeleteOneLeaveRequest(ctx *gin.Context) {
	id := ctx.Query("_id")

	err := l.LeaveRequestUseCase.DeleteOne(ctx, id)
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

package leave_request_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpdateStatusApproved update the leave request's information
// @Summary Update Leave Request Information
// @Description Updates the leave request's information
// @Tags Leave Request
// @Accept json
// @Produce json
// @Router /api/v1/leave-requests/update/_id [put]
// @Security CookieAuth
func (l *LeaveRequestController) UpdateStatusApproved(ctx *gin.Context) {
	_id := ctx.Query("_id")

	err := l.LeaveRequestUseCase.UpdateOneWithApproved(ctx, _id)
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

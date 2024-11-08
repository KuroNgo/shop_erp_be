package leave_request_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	leaverequestdomain "shop_erp_mono/internal/domain/human_resource_management/leave_request"
)

// UpdateOne update the leave request's information
// @Summary Update Leave Request Information
// @Description Updates the leave request's information
// @Tags Leave Request
// @Accept json
// @Produce json
// @Param LeaveRequest body leave_request_domain.Input true "Leave Request data"
// @Router /api/v1/leave-requests/update [put]
// @Security CookieAuth
func (l *LeaveRequestController) UpdateOne(ctx *gin.Context) {
	var input leaverequestdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	_id := ctx.Query("_id")

	err := l.LeaveRequestUseCase.UpdateOne(ctx, _id, &input)
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

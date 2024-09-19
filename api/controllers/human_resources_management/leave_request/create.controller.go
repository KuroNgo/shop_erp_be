package leave_request_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	leaverequestdomain "shop_erp_mono/domain/human_resource_management/leave_request"
)

// CreateOne create the leave request's information
// @Summary Create Leave Request Information
// @Description Create the leave request's information
// @Tags Leave Request
// @Accept json
// @Produce json
// @Param LeaveRequest body leave_request_domain.Input true "Leave Request data"
// @Router /api/v1/leave_requests/create [post]
// @Security CookieAuth
func (l *LeaveRequestController) CreateOne(ctx *gin.Context) {
	var input leaverequestdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	err := l.LeaveRequestUseCase.CreateOne(ctx, &input)
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

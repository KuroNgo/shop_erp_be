package leave_request_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAll get all the leave request's information
// @Summary Get all Leave Request Information
// @Description Get all the leave request's information
// @Tags Leave Request
// @Accept json
// @Produce json
// @Router /api/v1/leave-requests/get/all [get]
// @Security CookieAuth
func (l *LeaveRequestController) GetAll(ctx *gin.Context) {
	data, err := l.LeaveRequestUseCase.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
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

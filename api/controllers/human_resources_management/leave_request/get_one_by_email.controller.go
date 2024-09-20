package leave_request_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByEmailEmployee get one by email the leave request's information
// @Summary Get one by email Leave Request Information
// @Description Get one by email the leave request's information
// @Tags Leave Request
// @Accept json
// @Produce json
// @Param email path string true "Email"
// @Router /api/v1/leave-requests/get/email [get]
// @Security CookieAuth
func (l *LeaveRequestController) GetByEmailEmployee(ctx *gin.Context) {
	email := ctx.Query("email")

	data, err := l.LeaveRequestUseCase.GetByEmailEmployee(ctx, email)
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

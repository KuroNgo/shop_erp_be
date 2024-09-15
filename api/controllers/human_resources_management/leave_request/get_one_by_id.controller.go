package leave_request_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// FetchOneByIDLeaveRequest get one by id the leave request's information
// @Summary Get one by id Leave Request Information
// @Description Get one by email the leave request's information
// @Tags Leave Request
// @Accept json
// @Produce json
// @Param _id path string true "Employee ID"
// @Router /api/v1/leave_requests/get/_id [get]
// @Security CookieAuth
func (l *LeaveRequestController) FetchOneByIDLeaveRequest(ctx *gin.Context) {
	id := ctx.Param("_id")

	data, err := l.LeaveRequestUseCase.GetOneByID(ctx, id)
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

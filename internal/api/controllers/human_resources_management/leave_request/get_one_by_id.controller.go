package leave_request_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByID get one by id the leave request's information
// @Summary Get one by id Leave Request Information
// @Description Get one by email the leave request's information
// @Tags Leave Request
// @Accept json
// @Produce json
// @Param _id query string true "Employee ID"
// @Router /api/v1/leave-requests/get/_id [get]
// @Security CookieAuth
func (l *LeaveRequestController) GetByID(ctx *gin.Context) {
	id := ctx.Query("_id")

	data, err := l.LeaveRequestUseCase.GetByID(ctx, id)
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

package log_activity_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAll retrieves the activity log's information
// @Summary Get Activity Log Information
// @Description Retrieves the activity log's information
// @Tags Log
// @Accept  json
// @Produce  json
// @Router /api/v1/activity-log/get/all [get]
// @Security CookieAuth
func (a *ActivityController) GetAll(ctx *gin.Context) {
	data, err := a.ActivityUseCase.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
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

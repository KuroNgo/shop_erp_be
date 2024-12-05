package helper_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetJobCount retrieves number of the cron
// @Summary Get Number Of the Cron
// @Description Retrieves Number Of the Cron
// @Tags Cronjob
// @Accept  json
// @Produce  json
// @Router /api/v1/helper/cron/count [get]
// @Security CookieAuth
func (h *HelperController) GetJobCount(ctx *gin.Context) {
	count := h.Cr.GetJobCount()
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"count":  count,
	})
}

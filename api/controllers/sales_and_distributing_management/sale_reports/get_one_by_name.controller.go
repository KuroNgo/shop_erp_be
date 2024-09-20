package sales_report_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetBySummary godoc
// @Summary Get Sales Report Summary by Date Range
// @Description Retrieve a summary of sales reports for a specified date range
// @Tags SalesReports
// @Accept json
// @Produce json
// @Param start_date query string true "Start Date"
// @Param end_date query string true "End Date"
// @Success 200 {object} map[string]interface{} "status: success, data: Retrieved Sales Report Summary"
// @Failure 400 {object} map[string]interface{} "status: error, message: Retrieval error"
// @Router /api/v1/sales-reports/get/summary [get]
func (s *SalesReportController) GetBySummary(ctx *gin.Context) {
	startDate := ctx.Query("start_date")
	endDate := ctx.Query("end_date")

	data, err := s.SalesReportUseCase.GetReportSummary(ctx, startDate, endDate)
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

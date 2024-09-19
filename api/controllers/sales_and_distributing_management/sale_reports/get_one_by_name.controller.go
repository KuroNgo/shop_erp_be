package sales_report_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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

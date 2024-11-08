package sales_report_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByDate godoc
// @Summary Get Sales Reports by Date
// @Description Retrieve all sales reports associated with a specific report date
// @Tags SalesReports
// @Accept json
// @Produce json
// @Param report_date query string true "Report Date"
// @Router /api/v1/sales-reports/get/report_date [get]
func (s *SalesReportController) GetByDate(ctx *gin.Context) {
	reportDate := ctx.Query("report_date")

	data, err := s.SalesReportUseCase.GetByDate(ctx, reportDate)
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

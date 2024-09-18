package sales_report_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *SalesReportController) GetByDate(ctx *gin.Context) {
	reportDate := ctx.Param("report_date")

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

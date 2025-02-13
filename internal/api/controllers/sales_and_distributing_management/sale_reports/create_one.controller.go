package sales_report_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	salereportsdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/sale_reports"
)

// CreateOne godoc
// @Summary Create a new Sales Report
// @Description This API creates a new sales report from the input data
// @Tags SalesReports
// @Accept json
// @Produce json
// @Param salesReport body sale_reports_domain.Input true "Sales Report information"
// @Router /api/v1/sales-reports/create [post]
func (s *SalesReportController) CreateOne(ctx *gin.Context) {
	var input salereportsdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if err := s.SalesReportUseCase.CreateOne(ctx, &input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

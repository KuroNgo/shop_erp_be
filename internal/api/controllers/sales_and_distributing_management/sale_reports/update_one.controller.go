package sales_report_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	salereportsdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/sale_reports"
)

// UpdateOne godoc
// @Summary Update a Sales Report
// @Description This API updates an existing sales report based on the provided ID and input data
// @Tags SalesReports
// @Accept json
// @Produce json
// @Param _id query string true "Sales Report ID"
// @Param salesReport body sale_reports_domain.Input true "Updated Sales Report information"
// @Router /api/v1/sales-reports/update [put]
func (s *SalesReportController) UpdateOne(ctx *gin.Context) {
	var input salereportsdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	_id := ctx.Query("_id")

	if err := s.SalesReportUseCase.UpdateOne(ctx, _id, &input); err != nil {
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

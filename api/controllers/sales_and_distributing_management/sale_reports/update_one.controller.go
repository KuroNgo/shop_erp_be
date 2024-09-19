package sales_report_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	salereportsdomain "shop_erp_mono/domain/sales_and_distribution_management/sale_reports"
)

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

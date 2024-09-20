package sales_report_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOne godoc
// @Summary Delete a Sales Report
// @Description This API deletes a sales report based on the provided ID
// @Tags SalesReports
// @Accept json
// @Produce json
// @Param _id query string true "Sales Report ID"
// @Router /api/v1/sales-reports/delete [delete]
func (s *SalesReportController) DeleteOne(ctx *gin.Context) {
	_id := ctx.Query("_id")

	if err := s.SalesReportUseCase.DeleteOne(ctx, _id); err != nil {
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

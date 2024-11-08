package sales_report_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByID godoc
// @Summary Get a Sales Report by ID
// @Description Retrieve a sales report based on the provided ID
// @Tags SalesReports
// @Accept json
// @Produce json
// @Param _id query string true "Sales Report ID"
// @Router /api/v1/sales-reports/_id [get]
func (s *SalesReportController) GetByID(ctx *gin.Context) {
	_id := ctx.Query("_id")

	data, err := s.SalesReportUseCase.GetByID(ctx, _id)
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

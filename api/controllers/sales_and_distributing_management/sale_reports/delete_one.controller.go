package sales_report_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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

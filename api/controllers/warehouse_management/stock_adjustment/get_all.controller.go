package stock_adjustment_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_erp_mono/repository"
	"strconv"
)

func (s *StockAdjustmentController) GetAll(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	pageValue, err := strconv.ParseInt(page, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	var paginate repository.Pagination
	paginate.Page = pageValue

	data, err := s.StockAdjustmentUseCase.GetAllWithPagination(ctx, paginate)
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

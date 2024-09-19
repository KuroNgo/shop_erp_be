package stockmovement_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_erp_mono/repository"
)

func (s *StockMovementController) GetAllPagination(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	var paginate repository.Pagination
	paginate.Page = page

	data, err := s.StockMovementUseCase.GetAllWithPagination(ctx, paginate)
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

package supplier_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_erp_mono/repository"
	"strconv"
)

// GetAllSupplierWithPagination get all the supplier information with pagination
// @Summary get all supplier Information
// @Description get all the supplier's information
// @Tags Supplier
// @Accept json
// @Produce json
// @Param page query string true "Page number for pagination"
// @Router /api/v1/suppliers/get/all/pagination [get]
// @Security CookieAuth
func (s *SupplierController) GetAllSupplierWithPagination(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	number, err := strconv.ParseInt(page, 10, 64)

	var pagination repository.Pagination
	pagination.Page = number

	data, err := s.SupplierUseCase.GetSuppliersWithPagination(ctx, pagination)
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
package supplier_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	supplierdomain "shop_erp_mono/domain/warehouse_management/supplier"
)

// CreateSupplier create the supplier information
// @Summary Create supplier Information
// @Description Create the supplier's information
// @Tags Supplier
// @Accept json
// @Produce json
// @Param Supplier body supplier_domain.Input true "Supplier data"
// @Router /api/v1/suppliers/create [post]
// @Security CookieAuth
func (s *SupplierController) CreateSupplier(ctx *gin.Context) {
	var input supplierdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	err := s.SupplierUseCase.CreateOne(ctx, &input)
	if err != nil {
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

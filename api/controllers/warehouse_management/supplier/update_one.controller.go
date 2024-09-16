package supplier_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	supplierdomain "shop_erp_mono/domain/warehouse_management/supplier"
)

// UpdateSupplier update by name the supplier information with pagination
// @Summary update by name supplier Information
// @Description update by name the supplier's information
// @Tags Supplier
// @Accept json
// @Produce json
// @Param Supplier body supplier_domain.Input true "Supplier data"
// @Param _id path string true "supplier ID"
// @Router /api/v1/suppliers/update [get]
// @Security CookieAuth
func (s *SupplierController) UpdateSupplier(ctx *gin.Context) {
	var input supplierdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	_id := ctx.Param("_id")

	err := s.SupplierUseCase.UpdateSupplier(ctx, _id, &input)
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

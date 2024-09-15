package supplier_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByNameSupplier get by name the supplier information with pagination
// @Summary get by name supplier Information
// @Description get by name the supplier's information
// @Tags Supplier
// @Accept json
// @Produce json
// @Param name path string true "supplier name"
// @Router /api/v1/suppliers/get/name [get]
// @Security CookieAuth
func (s *SupplierController) GetByNameSupplier(ctx *gin.Context) {
	name := ctx.Param("name")

	data, err := s.SupplierUseCase.GetSupplierByID(ctx, name)
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

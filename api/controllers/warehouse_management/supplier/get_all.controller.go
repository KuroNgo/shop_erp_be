package supplier_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllSupplier get all the supplier information with pagination
// @Summary get all supplier Information
// @Description get all the supplier's information
// @Tags Supplier
// @Accept json
// @Produce json
// @Router /api/v1/suppliers/get/all [get]
// @Security CookieAuth
func (s *SupplierController) GetAllSupplier(ctx *gin.Context) {
	data, err := s.SupplierUseCase.GetSuppliers(ctx)
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

package supplier_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByIDSupplier get by id the supplier information with pagination
// @Summary get by id supplier Information
// @Description get by id the supplier's information
// @Tags Supplier
// @Accept json
// @Produce json
// @Param _id path string true "supplier ID"
// @Router /api/v1/suppliers/get/_id [get]
// @Security CookieAuth
func (s *SupplierController) GetByIDSupplier(ctx *gin.Context) {
	_id := ctx.Query("_id")

	data, err := s.SupplierUseCase.GetSupplierByID(ctx, _id)
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

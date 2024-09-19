package supplier_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteSupplier delete the supplier information
// @Summary delete supplier Information
// @Description delete the supplier's information
// @Tags Supplier
// @Accept json
// @Produce json
// @Param _id path string true "supplier ID"
// @Router /api/v1/suppliers/delete [delete]
// @Security CookieAuth
func (s *SupplierController) DeleteSupplier(ctx *gin.Context) {
	_id := ctx.Query("_id")

	err := s.SupplierUseCase.DeleteSupplier(ctx, _id)
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

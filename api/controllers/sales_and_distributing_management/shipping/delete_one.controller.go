package shipping_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOne godoc
// @Summary Delete a Shipping Entry
// @Description This API deletes a shipping entry based on the provided ID
// @Tags Shipping
// @Accept json
// @Produce json
// @Param _id query string true "Shipping Entry ID"
// @Router /api/v1/shipping/delete [delete]
func (s *ShippingController) DeleteOne(ctx *gin.Context) {
	_id := ctx.Query("_id")

	if err := s.ShippingUseCase.DeleteOne(ctx, _id); err != nil {
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

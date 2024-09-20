package shipping_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByStatus godoc
// @Summary Get Shipping Entries by Status
// @Description Retrieve shipping entries associated with a specific status
// @Tags Shipping
// @Accept json
// @Produce json
// @Param status query string true "Shipping Status"
// @Router /api/v1/shipping/get/status [get]
func (s *ShippingController) GetByStatus(ctx *gin.Context) {
	status := ctx.Query("status")

	data, err := s.ShippingUseCase.GetByStatus(ctx, status)
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

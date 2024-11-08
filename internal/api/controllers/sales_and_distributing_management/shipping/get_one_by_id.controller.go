package shipping_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByID godoc
// @Summary Get a Shipping Entry by ID
// @Description Retrieve a shipping entry based on the provided ID
// @Tags Shipping
// @Accept json
// @Produce json
// @Param id query string true "Shipping Entry ID"
// @Router /api/v1/shipping/_id [get]
func (s *ShippingController) GetByID(ctx *gin.Context) {
	_id := ctx.Query("_id")

	data, err := s.ShippingUseCase.GetByID(ctx, _id)
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

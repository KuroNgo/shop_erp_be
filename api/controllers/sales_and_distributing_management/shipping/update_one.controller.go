package shipping_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	shippingdomain "shop_erp_mono/domain/sales_and_distribution_management/shipping"
)

// UpdateOne godoc
// @Summary Update a Shipping Entry
// @Description Update an existing shipping entry based on the provided ID and input data
// @Tags Shipping
// @Accept json
// @Produce json
// @Param _id query string true "Shipping Entry ID"
// @Param shipping body shipping_domain.Input true "Updated Shipping information"
// @Router /api/v1/shipping/update [put]
func (s *ShippingController) UpdateOne(ctx *gin.Context) {
	var input shippingdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	_id := ctx.Query("_id")

	if err := s.ShippingUseCase.UpdateOne(ctx, _id, &input); err != nil {
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

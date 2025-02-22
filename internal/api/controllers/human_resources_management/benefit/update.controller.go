package benefit_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	benefitsdomain "shop_erp_mono/internal/domain/human_resource_management/benefits"
)

// UpdateOne updates the benefit's information
// @Summary Update Benefit Information
// @Description Updates the benefit's information
// @Tags Benefit
// @Produce json
// @Param _id query string true "Benefit ID"
// @Param attendance body benefits_domain.Input true "Benefit data"
// @Router /api/v1/benefits/_id [put]
// @Security CookieAuth
func (b *BenefitController) UpdateOne(ctx *gin.Context) {
	var input benefitsdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	attendanceID := ctx.Query("_id")

	if err := b.BenefitUseCase.UpdateOne(ctx, attendanceID, &input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

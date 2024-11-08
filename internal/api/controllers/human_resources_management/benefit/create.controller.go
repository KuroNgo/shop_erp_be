package benefit_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	benefitsdomain "shop_erp_mono/internal/domain/human_resource_management/benefits"
)

// CreateOne create the benefit's information
// @Summary Create Benefit Information
// @Description Create the benefit's information
// @Tags Benefit
// @Accept json
// @Produce json
// @Param Benefit body benefits_domain.Input true "Benefit data"
// @Router /api/v1/benefits/create [post]
// @Security CookieAuth
func (b *BenefitController) CreateOne(ctx *gin.Context) {
	var input benefitsdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	err := b.BenefitUseCase.CreateOne(ctx, &input)
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

package benefit_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	benefits_domain "shop_erp_mono/domain/human_resource_management/benefits"
)

// CreateOneBenefit create the benefit's information
// @Summary Create Benefit Information
// @Description Create the benefit's information
// @Tags Benefit
// @Accept json
// @Produce json
// @Param Benefit body benefits_domain.Input true "Benefit data"
// @Router /api/v1/benefits/create [post]
// @Security CookieAuth
func (b *BenefitController) CreateOneBenefit(ctx *gin.Context) {
	var input benefits_domain.Input
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

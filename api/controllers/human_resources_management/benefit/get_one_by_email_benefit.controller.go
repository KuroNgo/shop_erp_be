package benefit_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// FetchOneBenefitByEmail retrieves the benefit's information
// @Summary Get Benefit Information By ID
// @Description Retrieves the benefit's information name
// @Tags Benefit
// @Produce  json
// @Param email path string true "Benefit ID"
// @Router /api/v1/benefits/get/one/email [get]
// @Security CookieAuth
func (b *BenefitController) FetchOneBenefitByEmail(ctx *gin.Context) {
	email := ctx.Param("email")

	data, err := b.BenefitUseCase.GetOneByEmail(ctx, email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
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

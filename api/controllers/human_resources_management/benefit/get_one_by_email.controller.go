package benefit_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByEmail retrieves the benefit's information
// @Summary Get Benefit Information By ID
// @Description Retrieves the benefit's information name
// @Tags Benefit
// @Produce  json
// @Param email path string true "Benefit ID"
// @Router /api/v1/benefits/get/email [get]
// @Security CookieAuth
func (b *BenefitController) GetByEmail(ctx *gin.Context) {
	email := ctx.Query("email")

	data, err := b.BenefitUseCase.GetByEmail(ctx, email)
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

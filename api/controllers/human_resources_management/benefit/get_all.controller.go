package benefit_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAll retrieves the benefit's information
// @Summary Get Benefit Information
// @Description Retrieves the benefit's information
// @Tags Benefit
// @Accept  json
// @Produce  json
// @Router /api/v1/benefits/get/all [get]
// @Security CookieAuth
func (b *BenefitController) GetAll(ctx *gin.Context) {
	data, err := b.BenefitUseCase.GetAll(ctx)
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

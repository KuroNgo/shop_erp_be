package benefit_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// FetchOneBenefitByID retrieves the benefit's information
// @Summary Get Benefit Information By ID
// @Description Retrieves the benefit's information name
// @Tags Benefit
// @Accept  json
// @Produce  json
// @Param _id path string true "Benefit ID"
// @Router /api/v1/benefits/get/_id [get]
// @Security CookieAuth
func (b *BenefitController) FetchOneBenefitByID(ctx *gin.Context) {
	attendanceID := ctx.Param("_id")

	data, err := b.BenefitUseCase.GetOneByID(ctx, attendanceID)
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

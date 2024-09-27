package benefit_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByID retrieves the benefit's information
// @Summary Get Benefit Information By ID
// @Description Retrieves the benefit's information name
// @Tags Benefit
// @Accept  json
// @Produce  json
// @Param _id query string true "Benefit ID"
// @Router /api/v1/benefits/get/_id [get]
// @Security CookieAuth
func (b *BenefitController) GetByID(ctx *gin.Context) {
	attendanceID := ctx.Query("_id")

	data, err := b.BenefitUseCase.GetByID(ctx, attendanceID)
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

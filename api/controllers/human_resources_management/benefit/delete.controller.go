package benefit_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOneBenefit delete the benefit's information
// @Summary Delete Benefit Information
// @Description Deletes the benefit's information by ID
// @Tags Benefit
// @Produce json
// @Param _id path string true "Benefit ID"
// @Router /api/v1/benefits/_id [delete]
// @Security CookieAuth
func (b *BenefitController) DeleteOneBenefit(ctx *gin.Context) {
	attendanceID := ctx.Param("_id")

	if err := b.BenefitUseCase.DeleteOne(ctx, attendanceID); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

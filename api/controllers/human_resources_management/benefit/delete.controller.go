package benefit_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOne delete the benefit's information
// @Summary Delete Benefit Information
// @Description Deletes the benefit's information by ID
// @Tags Benefit
// @Produce json
// @Param _id query string true "Benefit ID"
// @Router /api/v1/benefits/_id [delete]
// @Security CookieAuth
func (b *BenefitController) DeleteOne(ctx *gin.Context) {
	attendanceID := ctx.Query("_id")

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

package budget_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteBudget delete the budget's information
// @Summary Delete Budget Information
// @Description Delete the budget's information
// @Tags Budget
// @Accept json
// @Produce json
// @Param _id path string true "Budget ID"
// @Router /api/v1/budgets/delete [delete]
// @Security CookieAuth
func (b BudgetController) DeleteOne(ctx *gin.Context) {
	_id := ctx.Query("_id")

	err := b.BudgetUseCase.DeleteOne(ctx, _id)
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

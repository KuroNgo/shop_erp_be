package budget_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	budgetsdomain "shop_erp_mono/domain/accounting_management/budgets"
)

// UpdateOne Get by name the budget's information
// @Summary Get by name Budget Information
// @Description Get by name the budget's information
// @Tags Budget
// @Accept json
// @Produce json
// @Param name path string true "Budget name"
// @Param Budget body budgets_domain.Input true "Budget data"
// @Router /api/v1/budgets/update [put]
// @Security CookieAuth
func (b BudgetController) UpdateOne(ctx *gin.Context) {
	var input budgetsdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	_id := ctx.Param("_id")

	err := b.BudgetUseCase.UpdateOne(ctx, _id, &input)
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

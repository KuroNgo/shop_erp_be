package budget_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	budgetsdomain "shop_erp_mono/domain/accounting_management/budgets"
)

// CreateBudget create the budget's information
// @Summary Create Budget Information
// @Description Create the budget's information
// @Tags Budget
// @Accept json
// @Produce json
// @Param Budget body budgets_domain.Input true "Budget data"
// @Router /api/v1/budgets/create [post]
// @Security CookieAuth
func (b BudgetController) CreateBudget(ctx *gin.Context) {
	var input budgetsdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	err := b.BudgetUseCase.CreateBudget(ctx, &input)
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

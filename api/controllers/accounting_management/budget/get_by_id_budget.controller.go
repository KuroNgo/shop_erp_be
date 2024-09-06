package budget_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByIDBudget Get by id the budget's information
// @Summary Get by id Budget Information
// @Description Get by id the budget's information
// @Tags Budget
// @Accept json
// @Produce json
// @Param _id path string true "Budget ID"
// @Router /api/v1/budgets/get/_id [get]
// @Security CookieAuth
func (b BudgetController) GetByIDBudget(ctx *gin.Context) {
	_id := ctx.Param("_id")

	data, err := b.BudgetUseCase.GetBudget(ctx, _id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
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

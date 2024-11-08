package budget_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByName Get by name the budget's information
// @Summary Get by name Budget Information
// @Description Get by name the budget's information
// @Tags Budget
// @Accept json
// @Produce json
// @Param name path string true "Budget name"
// @Router /api/v1/budgets/get/name [get]
// @Security CookieAuth
func (b BudgetController) GetByName(ctx *gin.Context) {
	name := ctx.Query("name")

	data, err := b.BudgetUseCase.GetByName(ctx, name)
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

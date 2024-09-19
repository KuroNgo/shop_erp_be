package budget_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAll Get all the budget's information
// @Summary Get Budget Information
// @Description Get the budget's information
// @Tags Budget
// @Accept json
// @Produce json
// @Router /api/v1/budgets/get/all [get]
// @Security CookieAuth
func (b BudgetController) GetAll(ctx *gin.Context) {
	data, err := b.BudgetUseCase.GetAll(ctx)
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

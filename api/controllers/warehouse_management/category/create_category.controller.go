package category_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	categorydomain "shop_erp_mono/domain/warehouse_management/product_category"
)

// CreateCategory Create a new category
// @Summary Create category
// @Description Create new category
// @Tags Category
// @Accept json
// @Produce json
// @Param Category body category_domain.Input true "Category data"
// @Security ApiKeyAuth
// @Router /api/v1/categories/create [post]
func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var input categorydomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if err := c.CategoryUseCase.CreateCategory(ctx, &input); err != nil {
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

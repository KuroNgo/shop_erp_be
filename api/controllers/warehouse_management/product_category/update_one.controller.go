package category_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	categorydomain "shop_erp_mono/domain/warehouse_management/product_category"
)

// UpdateCategory Update product_category
// @Summary Update product_category
// @Description Update product_category
// @Tags Category
// @Accept json
// @Produce json
// @Param Category body category_domain.Input true "Category data"
// @Param name path string true "Category ID"
// @Security ApiKeyAuth
// @Router /api/v1/categories/update [put]
func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	var input categorydomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	_id := ctx.Query("_id")

	if err := c.CategoryUseCase.UpdateOne(ctx, _id, &input); err != nil {
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

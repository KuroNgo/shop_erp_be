package category_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByNameCategories Get by name product_category
// @Summary Get by name product_category
// @Description Get by id product_category
// @Tags Category
// @Accept json
// @Produce json
// @Param name path string true "Category ID"
// @Security ApiKeyAuth
// @Router /api/v1/categories/get/name [get]
func (c *CategoryController) GetByNameCategories(ctx *gin.Context) {
	name := ctx.Query("name")

	data, err := c.CategoryUseCase.GetByName(ctx, name)
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

package category_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByNameCategories Get by name category
// @Summary Get by name category
// @Description Get by id category
// @Tags Category
// @Accept json
// @Produce json
// @Param name path string true "Category ID"
// @Security ApiKeyAuth
// @Router /api/v1/categories/get/name [get]
func (c *CategoryController) GetByNameCategories(ctx *gin.Context) {
	name := ctx.Param("name")

	data, err := c.CategoryUseCase.GetByNameCategory(ctx, name)
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
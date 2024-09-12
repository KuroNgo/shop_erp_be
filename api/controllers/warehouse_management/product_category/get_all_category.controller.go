package category_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllCategories Get all product_category
// @Summary Get all product_category
// @Description Get all product_category
// @Tags Category
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Router /api/v1/categories/get/all [get]
func (c *CategoryController) GetAllCategories(ctx *gin.Context) {
	data, err := c.CategoryUseCase.GetAllCategories(ctx)
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

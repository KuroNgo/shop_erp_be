package category_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByIDCategories Get by id product_category
// @Summary Get by id product_category
// @Description Get by id product_category
// @Tags Category
// @Accept json
// @Produce json
// @Param _id path string true "Category ID"
// @Security ApiKeyAuth
// @Router /api/v1/categories/get/_id [get]
func (c *CategoryController) GetByIDCategories(ctx *gin.Context) {
	_id := ctx.Query("_id")

	data, err := c.CategoryUseCase.GetByID(ctx, _id)
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

package category_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByIDCategories Get by id category
// @Summary Get by id category
// @Description Get by id category
// @Tags Category
// @Accept json
// @Produce json
// @Param _id path string true "Category ID"
// @Security ApiKeyAuth
// @Router /api/v1/categories/get/_id [get]
func (c *CategoryController) GetByIDCategories(ctx *gin.Context) {
	_id := ctx.Param("_id")

	data, err := c.CategoryUseCase.GetByIDCategory(ctx, _id)
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
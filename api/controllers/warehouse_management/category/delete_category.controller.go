package category_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteCategory Delete a new category
// @Summary Delete category
// @Description Delete new category
// @Tags Category
// @Accept json
// @Produce json
// @Param _id path string true "Category ID"
// @Security ApiKeyAuth
// @Router /api/v1/categories/delete [delete]
func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	_id := ctx.Param("_id")

	if err := c.CategoryUseCase.DeleteCategory(ctx, _id); err != nil {
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

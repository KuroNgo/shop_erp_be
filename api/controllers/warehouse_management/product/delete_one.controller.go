package product_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOneProduct delete the product's information
// @Summary Delete Product Information
// @Description Deletes the product's information
// @Tags Product
// @Accept json
// @Produce json
// @Param _id path string true "Product ID"
// @Router /api/v1/products/delete [delete]
// @Security CookieAuth
func (p *ProductController) DeleteOneProduct(ctx *gin.Context) {
	id := ctx.Query("_id")

	err := p.ProductUseCase.DeleteProduct(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

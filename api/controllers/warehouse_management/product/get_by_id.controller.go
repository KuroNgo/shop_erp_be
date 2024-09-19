package product_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetOneByIDProduct retrieves the product's information
// @Summary Get Product Information By ID
// @Description Retrieves the product's information id
// @Tags Product
// @Accept  json
// @Produce  json
// @Param _id path string true "Product ID"
// @Router /api/v1/products/get/_id [get]
// @Security CookieAuth
func (p *ProductController) GetOneByIDProduct(ctx *gin.Context) {
	id := ctx.Query("_id")

	data, err := p.ProductUseCase.GetProductByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
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

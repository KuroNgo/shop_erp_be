package product_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetOneByNameProduct retrieves the wm_product's information
// @Summary Get Product Information By Name
// @Description Retrieves the wm_product's information id
// @Tags Product
// @Accept  json
// @Produce  json
// @Param _id path string true "Product ID"
// @Router /api/v1/products/get/name [get]
// @Security CookieAuth
func (p *ProductController) GetOneByNameProduct(ctx *gin.Context) {
	name := ctx.Query("name")

	data, err := p.ProductUseCase.GetByName(ctx, name)
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

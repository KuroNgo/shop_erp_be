package product_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllProduct retrieves the product's information
// @Summary Get Product Information
// @Description Retrieves the product's information
// @Tags Product
// @Accept  json
// @Produce  json
// @Router /api/v1/products/get/all [get]
// @Security CookieAuth
func (p *ProductController) GetAllProduct(ctx *gin.Context) {
	data, err := p.ProductUseCase.GetAll(ctx)
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

package product_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	productdomain "shop_erp_mono/domain/warehouse_management/product"
)

// CreateProduct Create a new product
// @Summary Create product
// @Description Create new product
// @Tags Product
// @Accept json
// @Produce json
// @Param Product body product_domain.Input true "Product data"
// @Security ApiKeyAuth
// @Router /api/v1/products/create [post]
func (p *ProductController) CreateProduct(ctx *gin.Context) {
	var input productdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if err := p.ProductUseCase.CreateProduct(ctx, &input); err != nil {
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

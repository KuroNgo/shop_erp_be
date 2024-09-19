package product_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	productdomain "shop_erp_mono/domain/warehouse_management/product"
)

// UpdateProduct Update a new product
// @Summary Update product
// @Description Update new product
// @Tags Product
// @Accept json
// @Produce json
// @Param Product body product_domain.Input true "Product data"
// @Security ApiKeyAuth
// @Router /api/v1/products/update [put]
func (p *ProductController) UpdateProduct(ctx *gin.Context) {
	var input productdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	id := ctx.Query("_id")

	err := p.ProductUseCase.UpdateProduct(ctx, id, &input)
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

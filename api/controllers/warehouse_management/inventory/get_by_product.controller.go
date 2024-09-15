package inventory_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByProductID retrieves inventory details by product ID
// @Summary Get inventory by product ID
// @Description Retrieve inventory details using the product ID
// @Tags Inventory
// @Accept json
// @Produce json
// @Param product_id path string true "Product ID"
// @Router /api/v1/inventory/product/{product_id} [get]
func (i *InventoryController) GetByProductID(ctx *gin.Context) {
	productId := ctx.Param("product_id")

	data, err := i.InventoryUseCase.GetInventoryByProduct(ctx, productId)
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

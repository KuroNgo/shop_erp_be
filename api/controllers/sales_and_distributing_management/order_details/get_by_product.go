package order_detail_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByProductID godoc
// @Summary Get Order Details by Product ID
// @Description Retrieve all order details associated with a given Product ID
// @Tags OrderDetails
// @Accept json
// @Produce json
// @Param product_id query string true "Product ID"
// @Success 200 {object} map[string]interface{} "status: success, data: Retrieved Order Details"
// @Failure 400 {object} map[string]interface{} "status: error, message: Retrieval error"
// @Router /api/v1/order-details/get/product_id [get]
func (o *OrderDetailController) GetByProductID(ctx *gin.Context) {
	productId := ctx.Query("product_id")

	data, err := o.OrderDetailUseCase.GetByProductID(ctx, productId)
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

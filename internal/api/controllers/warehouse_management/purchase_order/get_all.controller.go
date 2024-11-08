package purchase_order_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_erp_mono/internal/repository"
)

// GetAll retrieves all purchase orders with pagination
// @Summary Get all purchase orders with pagination
// @Description Retrieve a list of purchase orders with pagination
// @Tags PurchaseOrder
// @Accept json
// @Produce json
// @Param page query string false "Page number" default(1)
// @Router /api/v1/purchase_orders/get/all [get]
func (p *PurchaseOrderController) GetAll(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	var paginate repository.Pagination
	paginate.Page = page

	data, err := p.PurchaseOrderUseCase.GetAllWithPagination(ctx, paginate)
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

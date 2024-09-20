package purchase_order_detail_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_erp_mono/repository"
)

// GetAllPagination godoc
// @Summary Get all purchase order details with pagination
// @Description Retrieve all purchase order details from the system with pagination support
// @Tags PurchaseOrderDetail
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Router /api/v1/purchase-order-details/pagination [get]
func (p *PurchaseOrderDetailController) GetAllPagination(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	var paginate repository.Pagination
	paginate.Page = page

	data, err := p.PurchaseOrderDetailUseCase.GetAllWithPagination(ctx, paginate)
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

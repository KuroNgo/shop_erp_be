package purchase_order_detail_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	purchase_order_detail_controller "shop_erp_mono/api/controllers/warehouse_management/purchase_order_detail"
	"shop_erp_mono/bootstrap"
	purchase_order_detail_domain "shop_erp_mono/domain/warehouse_management/purchase_order_detail"
	purchase_order_detail_repository "shop_erp_mono/repository/warehouse_management/purchase_order_detail/repository"
	purchase_order_detail_usecase "shop_erp_mono/usecase/warehouse_management/purchase_order_detail/usecase"
	"time"
)

func PurchaseOrderDetailRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	po := purchase_order_detail_repository.NewPurchaseOrderDetailRepository(db, purchase_order_detail_domain.CollectionPurchaseOrderDetail)
	purchaseOrderDetail := &purchase_order_detail_controller.PurchaseOrderDetailController{
		PurchaseOrderDetailUseCase: purchase_order_detail_usecase.NewProductOrderDetailRepository(timeout, po),
		Database:                   env,
	}

	router := group.Group("/purchase_order_details")
	router.GET("/get/_id", purchaseOrderDetail.GetByID)
	router.GET("/get/purchase_order_id", purchaseOrderDetail.GetByIPurchaseOrderD)
	router.GET("/get/all", purchaseOrderDetail.GetAll)
	router.POST("/create", purchaseOrderDetail.CreateOne)
	router.PUT("/update", purchaseOrderDetail.UpdateOne)
	router.DELETE("/delete", purchaseOrderDetail.DeleteOne)
}

package purchase_order_detail_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	purchaseorderdetailcontroller "shop_erp_mono/internal/api/controllers/warehouse_management/purchase_order_detail"
	"shop_erp_mono/internal/config"
	productdomain "shop_erp_mono/internal/domain/warehouse_management/product"
	purchaseorderdomain "shop_erp_mono/internal/domain/warehouse_management/purchase_order"
	purchaseorderdetaildomain "shop_erp_mono/internal/domain/warehouse_management/purchase_order_detail"
	purchaseorderrepository "shop_erp_mono/internal/repository/warehouse_management/purchase_order/repository"
	purchaseorderdetailrepository "shop_erp_mono/internal/repository/warehouse_management/purchase_order_detail/repository"
	productrepository "shop_erp_mono/internal/repository/warehouse_management/wm_product/repository"
	purchaseorderdetailusecase "shop_erp_mono/internal/usecase/warehouse_management/purchase_order_detail/usecase"
	"time"
)

func PurchaseOrderDetailRouter(env *config.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup, cacheTTL time.Duration) {
	po := purchaseorderdetailrepository.NewPurchaseOrderDetailRepository(db, purchaseorderdetaildomain.CollectionPurchaseOrderDetail)
	p := purchaseorderrepository.NewPurchaseOrderRepository(db, purchaseorderdomain.CollectionPurchaseOrder)
	pr := productrepository.NewProductRepository(db, productdomain.CollectionProduct)
	purchaseOrderDetail := &purchaseorderdetailcontroller.PurchaseOrderDetailController{
		PurchaseOrderDetailUseCase: purchaseorderdetailusecase.NewProductOrderDetailRepository(timeout, po, p, pr, cacheTTL),
		Database:                   env,
	}

	router := group.Group("/purchase-order-details")
	router.GET("/get/_id", purchaseOrderDetail.GetByID)
	router.GET("/get/purchase_order_id", purchaseOrderDetail.GetByIPurchaseOrderD)
	router.GET("/get/all", purchaseOrderDetail.GetAll)
	router.POST("/create", purchaseOrderDetail.CreateOne)
	router.PUT("/update", purchaseOrderDetail.UpdateOne)
	router.DELETE("/delete", purchaseOrderDetail.DeleteOne)
}

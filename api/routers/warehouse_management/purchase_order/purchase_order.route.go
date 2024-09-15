package purchase_order_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	purchase_order_controller "shop_erp_mono/api/controllers/warehouse_management/purchase_order"
	"shop_erp_mono/bootstrap"
	purchase_order_domain "shop_erp_mono/domain/warehouse_management/purchase_order"
	supplierdomain "shop_erp_mono/domain/warehouse_management/supplier"
	purchase_order_repository "shop_erp_mono/repository/warehouse_management/purchase_order/repository"
	supplier_repository "shop_erp_mono/repository/warehouse_management/supplier/repository"
	purchase_order_usecase "shop_erp_mono/usecase/warehouse_management/purchase_order/usecase"
	"time"
)

func PurchaseOrderRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	po := purchase_order_repository.NewPurchaseOrderRepository(db, purchase_order_domain.CollectionPurchaseOrder)
	su := supplier_repository.NewSupplierRepository(db, supplierdomain.CollectionSupplier)
	purchaseOrder := &purchase_order_controller.PurchaseOrderController{
		PurchaseOrderUseCase: purchase_order_usecase.NewPurchaseOrderUseCase(timeout, po, su),
		Database:             env,
	}

	router := group.Group("/purchase_orders")
	router.GET("/get/_id", purchaseOrder.GetByID)
	router.GET("/get/supplier_id", purchaseOrder.GetBySupplierID)
	router.GET("/get/all", purchaseOrder.GetAll)
	router.POST("/create", purchaseOrder.CreateOne)
	router.PUT("/update", purchaseOrder.UpdateOne)
	router.DELETE("/delete", purchaseOrder.DeleteOne)
}

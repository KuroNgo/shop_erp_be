package purchase_order_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	purchaseordercontroller "shop_erp_mono/api/controllers/warehouse_management/purchase_order"
	"shop_erp_mono/bootstrap"
	purchaseorderdomain "shop_erp_mono/domain/warehouse_management/purchase_order"
	supplierdomain "shop_erp_mono/domain/warehouse_management/supplier"
	purchaseorderrepository "shop_erp_mono/repository/warehouse_management/purchase_order/repository"
	supplierrepository "shop_erp_mono/repository/warehouse_management/supplier/repository"
	purchaseorderusecase "shop_erp_mono/usecase/warehouse_management/purchase_order/usecase"
	"time"
)

func PurchaseOrderRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup, cacheTTL time.Duration) {
	po := purchaseorderrepository.NewPurchaseOrderRepository(db, purchaseorderdomain.CollectionPurchaseOrder)
	su := supplierrepository.NewSupplierRepository(db, supplierdomain.CollectionSupplier)
	purchaseOrder := &purchaseordercontroller.PurchaseOrderController{
		PurchaseOrderUseCase: purchaseorderusecase.NewPurchaseOrderUseCase(timeout, po, su, cacheTTL),
		Database:             env,
	}

	router := group.Group("/purchase-orders")
	router.GET("/get/_id", purchaseOrder.GetByID)
	router.GET("/get/supplier_id", purchaseOrder.GetBySupplierID)
	router.GET("/get/all", purchaseOrder.GetAll)
	router.POST("/create", purchaseOrder.CreateOne)
	router.PUT("/update", purchaseOrder.UpdateOne)
	router.DELETE("/delete", purchaseOrder.DeleteOne)
}

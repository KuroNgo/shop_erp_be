package sale_orders_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	sales_order_controller "shop_erp_mono/api/controllers/sales_and_distributing_management/sale_orders"
	"shop_erp_mono/bootstrap"
	sale_orders_domain "shop_erp_mono/domain/sales_and_distribution_management/sale_orders"
	sales_order_repository "shop_erp_mono/repository/sales_and_distribution_management/sale_order/repository"
	sales_order_usecase "shop_erp_mono/usecase/sales_and_distribution_management/sale_order/usecase"
	"time"
)

func SaleOrderRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	so := sales_order_repository.NewSaleOrderRepository(db, sale_orders_domain.CollectionSalesOrder)
	salesOrder := &sales_order_controller.SalesOrderController{
		SalesOrderUseCase: sales_order_usecase.NewSaleOrderUseCase(timeout, so),
		Database:          env,
	}

	router := group.Group("/sales_orders")
	router.GET("/get/_id", salesOrder.GetByID)
	router.GET("/get/status", salesOrder.GetByStatus)
	router.GET("/get/customer_id", salesOrder.GetByCustomerID)
	router.POST("/create", salesOrder.CreateOne)
	router.PUT("/update", salesOrder.UpdateOne)
	router.DELETE("/delete", salesOrder.DeleteOne)
}

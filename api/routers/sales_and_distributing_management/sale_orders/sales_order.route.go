package sale_orders_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	salesordercontroller "shop_erp_mono/api/controllers/sales_and_distributing_management/sale_orders"
	"shop_erp_mono/bootstrap"
	customerdomain "shop_erp_mono/domain/sales_and_distribution_management/customer"
	saleordersdomain "shop_erp_mono/domain/sales_and_distribution_management/sale_orders"
	customerrepository "shop_erp_mono/repository/sales_and_distribution_management/customer/repository"
	salesorderrepository "shop_erp_mono/repository/sales_and_distribution_management/sale_order/repository"
	salesorderusecase "shop_erp_mono/usecase/sales_and_distribution_management/sale_order/usecase"
	"time"
)

func SaleOrderRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	so := salesorderrepository.NewSaleOrderRepository(db, saleordersdomain.CollectionSalesOrder)
	cu := customerrepository.NewCustomerRepository(db, customerdomain.CollectionCustomer)
	salesOrder := &salesordercontroller.SalesOrderController{
		SalesOrderUseCase: salesorderusecase.NewSaleOrderUseCase(timeout, so, cu),
		Database:          env,
	}

	router := group.Group("/sales-orders")
	router.GET("/get/_id", salesOrder.GetByID)
	router.GET("/get/status", salesOrder.GetByStatus)
	router.GET("/get/customer_id", salesOrder.GetByCustomerID)
	router.POST("/create", salesOrder.CreateOne)
	router.PUT("/update", salesOrder.UpdateOne)
	router.DELETE("/delete", salesOrder.DeleteOne)
}

package invoices_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	invoicecontroller "shop_erp_mono/api/controllers/sales_and_distributing_management/invoices"
	"shop_erp_mono/bootstrap"
	invoicesdomain "shop_erp_mono/domain/sales_and_distribution_management/invoices"
	sale_orders_domain "shop_erp_mono/domain/sales_and_distribution_management/sale_orders"
	invoicerepository "shop_erp_mono/repository/sales_and_distribution_management/invoices/repository"
	sales_order_repository "shop_erp_mono/repository/sales_and_distribution_management/sale_order/repository"
	invoiceusecase "shop_erp_mono/usecase/sales_and_distribution_management/invoices/usecase"
	"time"
)

func InvoiceRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	in := invoicerepository.NewInvoiceRepository(db, invoicesdomain.CollectionInvoice)
	so := sales_order_repository.NewSaleOrderRepository(db, sale_orders_domain.CollectionSalesOrder)
	invoice := &invoicecontroller.InvoiceController{
		InvoiceUseCase: invoiceusecase.NewInvoiceUseCase(timeout, in, so),
		Database:       env,
	}

	router := group.Group("/invoices")
	router.GET("/get/_id", invoice.GetByID)
	router.GET("/get/name", invoice.GetByName)
	router.GET("/get/status", invoice.GetByStatus)
	router.POST("/create", invoice.CreateOne)
	router.PUT("/update", invoice.UpdateOne)
	router.DELETE("/delete", invoice.DeleteOne)
}

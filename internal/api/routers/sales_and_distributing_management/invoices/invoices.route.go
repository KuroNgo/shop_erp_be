package invoices_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	invoicecontroller "shop_erp_mono/internal/api/controllers/sales_and_distributing_management/invoices"
	"shop_erp_mono/internal/config"
	invoicesdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/invoices"
	saleordersdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/sale_orders"
	invoicerepository "shop_erp_mono/internal/repository/sales_and_distribution_management/invoices/repository"
	salesorderrepository "shop_erp_mono/internal/repository/sales_and_distribution_management/sale_order/repository"
	invoiceusecase "shop_erp_mono/internal/usecase/sales_and_distribution_management/invoices/usecase"
	"time"
)

func InvoiceRouter(env *config.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup, cacheTTL time.Duration) {
	in := invoicerepository.NewInvoiceRepository(db, invoicesdomain.CollectionInvoice)
	so := salesorderrepository.NewSaleOrderRepository(db, saleordersdomain.CollectionSalesOrder)
	invoice := &invoicecontroller.InvoiceController{
		InvoiceUseCase: invoiceusecase.NewInvoiceUseCase(timeout, in, so, cacheTTL),
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

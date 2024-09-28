package sales_and_distributing_management

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"shop_erp_mono/api/middlewares"
	customerroute "shop_erp_mono/api/routers/sales_and_distributing_management/customer"
	invoicesroute "shop_erp_mono/api/routers/sales_and_distributing_management/invoices"
	orderdetailsroute "shop_erp_mono/api/routers/sales_and_distributing_management/order_details"
	paymentsroute "shop_erp_mono/api/routers/sales_and_distributing_management/payments"
	saleordersroute "shop_erp_mono/api/routers/sales_and_distributing_management/sale_orders"
	salereportsroute "shop_erp_mono/api/routers/sales_and_distributing_management/sale_reports"
	shippingroute "shop_erp_mono/api/routers/sales_and_distributing_management/shipping"
	"shop_erp_mono/bootstrap"
	"time"
)

func SetUp(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, gin *gin.Engine, cacheTTL time.Duration) {
	publicRouter := gin.Group("/api/v1")

	// Middleware
	publicRouter.Use(
		middlewares.CORSPublic(),
		middlewares.Recover(),
		gzip.Gzip(gzip.DefaultCompression,
			gzip.WithExcludedPaths([]string{",*"})),
		//middlewares.StructuredLogger(&log.Logger, value),
	)

	// All Public APIs
	customerroute.CustomerRouter(env, timeout, db, publicRouter, cacheTTL)
	invoicesroute.InvoiceRouter(env, timeout, db, publicRouter, cacheTTL)
	orderdetailsroute.OrderDetailRouter(env, timeout, db, publicRouter, cacheTTL)
	paymentsroute.PaymentRouter(env, timeout, db, publicRouter, cacheTTL)
	saleordersroute.SaleOrderRouter(env, timeout, db, publicRouter, cacheTTL)
	salereportsroute.SaleReportRouter(env, timeout, db, publicRouter, cacheTTL)
	shippingroute.ShippingRouter(env, timeout, db, publicRouter, cacheTTL)
}

package sales_and_distributing_management

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"shop_erp_mono/internal/api/middlewares"
	"shop_erp_mono/internal/api/routers/log_activity"
	customerroute "shop_erp_mono/internal/api/routers/sales_and_distributing_management/customer"
	invoicesroute "shop_erp_mono/internal/api/routers/sales_and_distributing_management/invoices"
	orderdetailsroute "shop_erp_mono/internal/api/routers/sales_and_distributing_management/order_details"
	paymentsroute "shop_erp_mono/internal/api/routers/sales_and_distributing_management/payments"
	saleordersroute "shop_erp_mono/internal/api/routers/sales_and_distributing_management/sale_orders"
	salereportsroute "shop_erp_mono/internal/api/routers/sales_and_distributing_management/sale_reports"
	shippingroute "shop_erp_mono/internal/api/routers/sales_and_distributing_management/shipping"
	"shop_erp_mono/internal/config"
	"shop_erp_mono/pkg/interface/security/casbin/middlewares"
	"shop_erp_mono/pkg/interface/security/casbin/principle"
	"time"
)

func SetUp(env *config.Database, client *mongo.Client, timeout time.Duration, db *mongo.Database, gin *gin.Engine, cacheTTL time.Duration) {
	publicRouter := gin.Group("/api/v1")
	value := log_activity.Activity(env, client, timeout, db, cacheTTL)
	enforcer := principle.SetUp(env)

	// Middleware
	publicRouter.Use(
		middlewares.CORSPrivate(),
		middlewares.Recover(),
		gzip.Gzip(gzip.DefaultCompression,
			gzip.WithExcludedPaths([]string{",*"})),
		casbin.Authorize(enforcer),
		middlewares.DeserializeUser(),
		middlewares.StructuredLogger(&log.Logger, value),
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

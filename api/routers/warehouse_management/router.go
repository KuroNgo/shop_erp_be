package warehouse_management

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"shop_erp_mono/api/middlewares"
	inventory_route "shop_erp_mono/api/routers/warehouse_management/inventory"
	product_route "shop_erp_mono/api/routers/warehouse_management/product"
	product_category_route "shop_erp_mono/api/routers/warehouse_management/product_category"
	purchase_order_route "shop_erp_mono/api/routers/warehouse_management/purchase_order"
	purchase_order_detail_route "shop_erp_mono/api/routers/warehouse_management/purchase_order_detail"
	stock_adjustment_route "shop_erp_mono/api/routers/warehouse_management/stock_adjustment"
	stockmovement_route "shop_erp_mono/api/routers/warehouse_management/stockmovement"
	supplier_route "shop_erp_mono/api/routers/warehouse_management/supplier"
	warehouse_route "shop_erp_mono/api/routers/warehouse_management/warehouse"
	"shop_erp_mono/bootstrap"
	casbin "shop_erp_mono/pkg/casbin/middlewares"
	"shop_erp_mono/pkg/casbin/principle"
	"time"
)

func SetUp(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, gin *gin.Engine, cacheTTL time.Duration) {
	publicRouter := gin.Group("/api/v1")

	// Khởi tạo Casbin enforcer
	enforcer := principle.SetUp()

	// Middleware
	publicRouter.Use(
		middlewares.CORSPublic(),
		middlewares.Recover(),
		gzip.Gzip(gzip.DefaultCompression,
			gzip.WithExcludedPaths([]string{",*"})),
		casbin.Authorize(enforcer),
		//middlewares.StructuredLogger(&log.Logger, value),
	)

	// All Public APIs
	product_route.ProductRouter(env, timeout, db, publicRouter, cacheTTL)
	product_category_route.ProductCategoryRouter(env, timeout, db, publicRouter, cacheTTL)
	purchase_order_detail_route.PurchaseOrderDetailRouter(env, timeout, db, publicRouter, cacheTTL)
	supplier_route.SupplierRouter(env, timeout, db, publicRouter, cacheTTL)
	inventory_route.InventoryRouter(env, timeout, db, publicRouter, cacheTTL)
	purchase_order_route.PurchaseOrderRouter(env, timeout, db, publicRouter, cacheTTL)
	warehouse_route.WarehouseRouter(env, timeout, db, publicRouter, cacheTTL)
	stock_adjustment_route.StockAdjustmentRouter(env, timeout, db, publicRouter, cacheTTL)
	stockmovement_route.StockMovementRouter(env, timeout, db, publicRouter, cacheTTL)
}

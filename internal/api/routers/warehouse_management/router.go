package warehouse_management

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"shop_erp_mono/internal/api/middlewares"
	inventoryroute "shop_erp_mono/internal/api/routers/warehouse_management/inventory"
	productroute "shop_erp_mono/internal/api/routers/warehouse_management/product"
	productcategoryroute "shop_erp_mono/internal/api/routers/warehouse_management/product_category"
	purchaseorderroute "shop_erp_mono/internal/api/routers/warehouse_management/purchase_order"
	purchaseorderdetailroute "shop_erp_mono/internal/api/routers/warehouse_management/purchase_order_detail"
	stockadjustmentroute "shop_erp_mono/internal/api/routers/warehouse_management/stock_adjustment"
	stockmovementroute "shop_erp_mono/internal/api/routers/warehouse_management/stockmovement"
	supplierroute "shop_erp_mono/internal/api/routers/warehouse_management/supplier"
	warehouseroute "shop_erp_mono/internal/api/routers/warehouse_management/warehouse"
	"shop_erp_mono/internal/config"
	"time"
)

func SetUp(env *config.Database, timeout time.Duration, db *mongo.Database, gin *gin.Engine, cacheTTL time.Duration) {
	publicRouter := gin.Group("/api/v1")

	// Khởi tạo Casbin enforcer
	//enforcer := principle.SetUp(env)

	// Middleware
	publicRouter.Use(
		middlewares.CORSPrivate(),
		middlewares.Recover(),
		gzip.Gzip(gzip.DefaultCompression,
			gzip.WithExcludedPaths([]string{",*"})),
		//casbin.Authorize(enforcer),
		//middlewares.StructuredLogger(&log.Logger, value),
	)

	// All Public APIs
	productroute.ProductRouter(env, timeout, db, publicRouter, cacheTTL)
	productcategoryroute.ProductCategoryRouter(env, timeout, db, publicRouter, cacheTTL)
	purchaseorderdetailroute.PurchaseOrderDetailRouter(env, timeout, db, publicRouter, cacheTTL)
	supplierroute.SupplierRouter(env, timeout, db, publicRouter, cacheTTL)
	inventoryroute.InventoryRouter(env, timeout, db, publicRouter, cacheTTL)
	purchaseorderroute.PurchaseOrderRouter(env, timeout, db, publicRouter, cacheTTL)
	warehouseroute.WarehouseRouter(env, timeout, db, publicRouter, cacheTTL)
	stockadjustmentroute.StockAdjustmentRouter(env, timeout, db, publicRouter, cacheTTL)
	stockmovementroute.StockMovementRouter(env, timeout, db, publicRouter, cacheTTL)
}

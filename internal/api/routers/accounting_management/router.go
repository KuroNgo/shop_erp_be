package accounting_management

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"shop_erp_mono/internal/api/middlewares"
	account_route "shop_erp_mono/internal/api/routers/accounting_management/account"
	budget_route "shop_erp_mono/internal/api/routers/accounting_management/budget"
	"shop_erp_mono/internal/config"
	"time"
)

func SetUp(env *config.Database, timeout time.Duration, db *mongo.Database, gin *gin.Engine) {
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
	account_route.AccountRouter(env, timeout, db, publicRouter)
	budget_route.BudgetRouter(env, timeout, db, publicRouter)
}

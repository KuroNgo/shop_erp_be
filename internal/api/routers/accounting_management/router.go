package accounting_management

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"shop_erp_mono/internal/api/middlewares"
	account_route "shop_erp_mono/internal/api/routers/accounting_management/account"
	budget_route "shop_erp_mono/internal/api/routers/accounting_management/budget"
	"shop_erp_mono/internal/api/routers/log_activity"
	"shop_erp_mono/internal/config"
	"shop_erp_mono/pkg/interface/security/casbin/middlewares"
	"shop_erp_mono/pkg/interface/security/casbin/principle"
	"time"
)

func SetUp(env *config.Database, client *mongo.Client, timeout time.Duration, db *mongo.Database, gin *gin.Engine, cacheTTL time.Duration) {
	publicRouter := gin.Group("/api/v1")
	value := log_activity.Activity(env, client, timeout, db, cacheTTL)

	// Khởi tạo Casbin enforcer
	enforcer := principle.SetUp(env)

	// Middleware
	publicRouter.Use(
		middlewares.CORSPrivate(),
		middlewares.Recover(),
		gzip.Gzip(gzip.DefaultCompression,
			gzip.WithExcludedPaths([]string{",*"})),
		casbin.Authorize(enforcer),
		middlewares.StructuredLogger(&log.Logger, value),
	)

	// All Public APIs
	account_route.AccountRouter(env, timeout, db, publicRouter)
	budget_route.BudgetRouter(env, timeout, db, publicRouter)
}

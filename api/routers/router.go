package routers

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"shop_erp_mono/api/middlewares"
	userroute "shop_erp_mono/api/routers/human_resources_management/user"
	swaggerroute "shop_erp_mono/api/routers/swagger"
	"shop_erp_mono/bootstrap"
	"time"
)

func SetUp(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, gin *gin.Engine) {
	SwaggerRouter := gin.Group("")
	publicRouter := gin.Group("/api")
	// Middleware
	publicRouter.Use(
		middlewares.CORSPublic(),
		//middleware.RateLimiter(),
		middlewares.Recover(),
		gzip.Gzip(gzip.DefaultCompression,
			gzip.WithExcludedPaths([]string{",*"})),
		//middlewares.StructuredLogger(&log.Logger, value),
	)

	// This is a CORS method for check IP validation
	publicRouter.OPTIONS("/*path", middlewares.OptionMessages)

	// All Public APIs
	userroute.UserRouter(env, timeout, db, publicRouter)
	swaggerroute.SwaggerRouter(env, timeout, db, SwaggerRouter)
}

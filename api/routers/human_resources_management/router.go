package human_resources_management

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"shop_erp_mono/api/middlewares"
	departmentroute "shop_erp_mono/api/routers/human_resources_management/department"
	roleroute "shop_erp_mono/api/routers/human_resources_management/role"
	userroute "shop_erp_mono/api/routers/human_resources_management/user"
	"shop_erp_mono/bootstrap"
	"time"
)

func SetUp(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("/api/v1")

	// Middleware
	publicRouter.Use(
		middlewares.CORSPublic(),
		middlewares.Recover(),
		gzip.Gzip(gzip.DefaultCompression,
			gzip.WithExcludedPaths([]string{",*"})),
		//middlewares.StructuredLogger(&log.Logger, value),
	)

	// This is a CORS method for check IP validation
	publicRouter.OPTIONS("/*path", middlewares.OptionMessages)

	// All Public APIs
	userroute.UserRouter(env, timeout, db, publicRouter)
	roleroute.RoleRouter(env, timeout, db, publicRouter)
	departmentroute.DepartmentRouter(env, timeout, db, publicRouter)
}
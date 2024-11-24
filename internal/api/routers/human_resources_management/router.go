package human_resources_management

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"shop_erp_mono/internal/api/middlewares"
	attendanceroute "shop_erp_mono/internal/api/routers/human_resources_management/attendance"
	benefitroute "shop_erp_mono/internal/api/routers/human_resources_management/benefit"
	candidateroute "shop_erp_mono/internal/api/routers/human_resources_management/candidate"
	contractroute "shop_erp_mono/internal/api/routers/human_resources_management/contract"
	departmentroute "shop_erp_mono/internal/api/routers/human_resources_management/department"
	employeeroute "shop_erp_mono/internal/api/routers/human_resources_management/employee"
	leaverequestroute "shop_erp_mono/internal/api/routers/human_resources_management/leave_request"
	performancereviewroute "shop_erp_mono/internal/api/routers/human_resources_management/performance_review"
	roleroute "shop_erp_mono/internal/api/routers/human_resources_management/role"
	salaryroute "shop_erp_mono/internal/api/routers/human_resources_management/salary"
	userroute "shop_erp_mono/internal/api/routers/human_resources_management/user"
	"shop_erp_mono/internal/config"
	casbin "shop_erp_mono/pkg/interface/casbin/middlewares"
	"shop_erp_mono/pkg/interface/casbin/principle"
	cronjob "shop_erp_mono/pkg/interface/cron"
	"time"
)

func SetUp(env *config.Database, cr *cronjob.CronScheduler, timeout time.Duration, db *mongo.Database, client *mongo.Client, gin *gin.Engine, cacheTTL time.Duration) {
	publicRouterV1 := gin.Group("/api/v1")
	publicRouterV2 := gin.Group("/api/v2")
	publicRouter := gin.Group("/api")
	router := gin.Group("")

	// Khởi tạo Casbin enforcer
	enforcer := principle.SetUp(env)

	// Middleware
	publicRouter.Use(
		middlewares.CORSPrivate(),
		middlewares.Recover(),
		gzip.Gzip(gzip.DefaultCompression,
			gzip.WithExcludedPaths([]string{",*"})),
		casbin.Authorize(enforcer),
		middlewares.DeserializeUser(),
		//middlewares.StructuredLogger(&log.Logger, value),
	)

	publicRouterV1.Use(
		middlewares.CORSPrivate(),
		middlewares.Recover(),
		gzip.Gzip(gzip.DefaultCompression,
			gzip.WithExcludedPaths([]string{",*"})),
		casbin.Authorize(enforcer),
		middlewares.DeserializeUser(),
		//middlewares.StructuredLogger(&log.Logger, value),
	)

	publicRouterV2.Use(
		middlewares.CORSPrivate(),
		middlewares.Recover(),
		gzip.Gzip(gzip.DefaultCompression,
			gzip.WithExcludedPaths([]string{",*"})),
		casbin.Authorize(enforcer),
		middlewares.DeserializeUser(),
		//middlewares.StructuredLogger(&log.Logger, value),
	)

	// This is a CORS method for check IP validation
	router.OPTIONS("/*path", middlewares.OptionMessages)

	// All Public APIs v1
	userroute.UserRouter(env, timeout, db, client, publicRouterV1)
	roleroute.RoleRouter(env, timeout, db, publicRouterV1, cacheTTL)
	departmentroute.DepartmentRouter(env, timeout, db, client, publicRouterV1, cacheTTL)
	salaryroute.SalaryRouter(env, timeout, db, publicRouterV1, cacheTTL)
	attendanceroute.AttendanceRouter(env, timeout, db, publicRouterV1, cacheTTL)
	employeeroute.EmployeeRouter(env, timeout, db, publicRouterV1, cacheTTL)
	benefitroute.BenefitRouter(env, timeout, db, publicRouterV1, cacheTTL)
	contractroute.ContractRouter(env, timeout, db, publicRouterV1, cacheTTL)
	leaverequestroute.LeaveRequestRouter(env, cr, timeout, db, client, publicRouterV1, cacheTTL)
	performancereviewroute.PerformanceReviewRouterV1(env, timeout, db, publicRouterV1, cacheTTL)
	candidateroute.CandidateRouter(env, timeout, db, client, publicRouter, cacheTTL)

	// All Public APIs v2
	performancereviewroute.PerformanceReviewRouterV2(env, timeout, db, publicRouterV2, cacheTTL)

	// All Public APIs
	userroute.UserGoogleRouter(env, timeout, db, client, publicRouter)
}

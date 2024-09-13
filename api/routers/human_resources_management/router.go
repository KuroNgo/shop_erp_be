package human_resources_management

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"shop_erp_mono/api/middlewares"
	attendanceroute "shop_erp_mono/api/routers/human_resources_management/attendance"
	benefitroute "shop_erp_mono/api/routers/human_resources_management/benefit"
	contractroute "shop_erp_mono/api/routers/human_resources_management/contract"
	departmentroute "shop_erp_mono/api/routers/human_resources_management/department"
	employeeroute "shop_erp_mono/api/routers/human_resources_management/employee"
	leaverequestroute "shop_erp_mono/api/routers/human_resources_management/leave_request"
	performancereviewroute "shop_erp_mono/api/routers/human_resources_management/performance_review"
	roleroute "shop_erp_mono/api/routers/human_resources_management/role"
	salaryroute "shop_erp_mono/api/routers/human_resources_management/salary"
	userroute "shop_erp_mono/api/routers/human_resources_management/user"
	"shop_erp_mono/bootstrap"
	"time"
)

func SetUp(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, gin *gin.Engine) {
	publicRouterV1 := gin.Group("/api/v1")
	publicRouterV2 := gin.Group("/api/v2")
	publicRouter := gin.Group("/api")
	router := gin.Group("")

	// Middleware
	publicRouterV1.Use(
		middlewares.CORSPublic(),
		middlewares.Recover(),
		gzip.Gzip(gzip.DefaultCompression,
			gzip.WithExcludedPaths([]string{",*"})),
		//middlewares.StructuredLogger(&log.Logger, value),
	)

	// This is a CORS method for check IP validation
	router.OPTIONS("/*path", middlewares.OptionMessages)

	// All Public APIs v1
	userroute.UserRouter(env, timeout, db, publicRouterV1)
	roleroute.RoleRouter(env, timeout, db, publicRouterV1)
	departmentroute.DepartmentRouter(env, timeout, db, publicRouterV1)
	salaryroute.SalaryRouter(env, timeout, db, publicRouterV1)
	attendanceroute.AttendanceRouter(env, timeout, db, publicRouterV1)
	employeeroute.EmployeeRouter(env, timeout, db, publicRouterV1)
	benefitroute.BenefitRouter(env, timeout, db, publicRouterV1)
	contractroute.ContractRouter(env, timeout, db, publicRouterV1)
	leaverequestroute.LeaveRequestRouter(env, timeout, db, publicRouterV1)
	performancereviewroute.PerformanceReviewRouterV1(env, timeout, db, publicRouterV1)

	// All Public APIs v2
	performancereviewroute.PerformanceReviewRouterV2(env, timeout, db, publicRouterV2)

	// All Public APIs
	userroute.UserGoogleRouter(env, timeout, db, publicRouter)
}

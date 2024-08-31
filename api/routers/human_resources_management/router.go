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
	leave_request_route "shop_erp_mono/api/routers/human_resources_management/leave_request"
	roleroute "shop_erp_mono/api/routers/human_resources_management/role"
	salaryroute "shop_erp_mono/api/routers/human_resources_management/salary"
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
	salaryroute.SalaryRouter(env, timeout, db, publicRouter)
	attendanceroute.AttendanceRouter(env, timeout, db, publicRouter)
	employeeroute.EmployeeRouter(env, timeout, db, publicRouter)
	benefitroute.BenefitRouter(env, timeout, db, publicRouter)
	contractroute.ContractRouter(env, timeout, db, publicRouter)
	leave_request_route.LeaveRequestRouter(env, timeout, db, publicRouter)
}

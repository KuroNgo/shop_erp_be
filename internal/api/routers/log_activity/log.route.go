package log_activity

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	logactivitycontroller "shop_erp_mono/internal/api/controllers/log_activity"
	"shop_erp_mono/internal/config"
	activitylogdomain "shop_erp_mono/internal/domain/activity_log"
	employeesdomain "shop_erp_mono/internal/domain/human_resource_management/employees"
	userdomain "shop_erp_mono/internal/domain/human_resource_management/user"
	logrepository "shop_erp_mono/internal/repository/activity_log/repository"
	employeerepository "shop_erp_mono/internal/repository/human_resource_management/employee/repository"
	userrepository "shop_erp_mono/internal/repository/human_resource_management/user/repository"
	activitylogusecase "shop_erp_mono/internal/usecase/activity_log/usecase"
	userusecase "shop_erp_mono/internal/usecase/human_resource_management/user/usecase"
	"time"
)

func Activity(env *config.Database, client *mongo.Client, timeout time.Duration, db *mongo.Database, cacheTTL time.Duration) *logactivitycontroller.ActivityController {
	ac := logrepository.NewLogRepository(activitylogdomain.CollectionActivityLog, db)
	users := userrepository.NewUserRepository(db, userdomain.CollectionUser)
	em := employeerepository.NewEmployeeRepository(db, employeesdomain.CollectionEmployee)

	activity := &logactivitycontroller.ActivityController{
		ActivityUseCase: activitylogusecase.NewActivityLogUseCase(timeout, ac, em, users, cacheTTL, client),
		UserUseCase:     userusecase.NewUserUseCase(env, timeout, users, client),
		Database:        env,
	}

	return activity
}

func ActivityRoute(env *config.Database, timeout time.Duration, db *mongo.Database, client *mongo.Client, group *gin.RouterGroup, cacheTTL time.Duration) {
	ac := logrepository.NewLogRepository(activitylogdomain.CollectionActivityLog, db)
	users := userrepository.NewUserRepository(db, userdomain.CollectionUser)
	em := employeerepository.NewEmployeeRepository(db, employeesdomain.CollectionEmployee)

	activity := &logactivitycontroller.ActivityController{
		ActivityUseCase: activitylogusecase.NewActivityLogUseCase(timeout, ac, em, users, cacheTTL, client),
		UserUseCase:     userusecase.NewUserUseCase(env, timeout, users, client),
		Database:        env,
	}

	router := group.Group("/activity")
	router.GET("/fetch", activity.GetAll)
}

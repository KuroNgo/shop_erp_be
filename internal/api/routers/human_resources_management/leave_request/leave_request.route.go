package leave_request_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	leaverequestcontroller "shop_erp_mono/internal/api/controllers/human_resources_management/leave_request"
	"shop_erp_mono/internal/config"
	departments_domain "shop_erp_mono/internal/domain/human_resource_management/departments"
	employees_domain "shop_erp_mono/internal/domain/human_resource_management/employees"
	leaverequestdomain "shop_erp_mono/internal/domain/human_resource_management/leave_request"
	departmentrepository "shop_erp_mono/internal/repository/human_resource_management/department/repository"
	employeerepository "shop_erp_mono/internal/repository/human_resource_management/employee/repository"
	leaverequestrepository "shop_erp_mono/internal/repository/human_resource_management/leave_request/repository"
	leave_request_usecase "shop_erp_mono/internal/usecase/human_resource_management/leave_request/usecase"
	cronjob "shop_erp_mono/pkg/interface/cron"
	"time"
)

func LeaveRequestRouter(env *config.Database, cr *cronjob.CronScheduler, timeout time.Duration, db *mongo.Database, client *mongo.Client, group *gin.RouterGroup, cacheTTL time.Duration) {
	lr := leaverequestrepository.NewLeaveRequestRepository(db, leaverequestdomain.CollectionLeaveRequest)
	em := employeerepository.NewEmployeeRepository(db, employees_domain.CollectionEmployee)
	de := departmentrepository.NewDepartmentRepository(db, departments_domain.CollectionDepartment)
	leaveRequest := &leaverequestcontroller.LeaveRequestController{
		LeaveRequestUseCase: leave_request_usecase.NewLeaveRequestUseCase(timeout, lr, em, cacheTTL, de, client),
		Database:            env,
		CronJob:             cr,
	}

	router := group.Group("/leave-requests")
	router.GET("/get/_id", leaveRequest.GetByID)
	router.GET("/get/email", leaveRequest.GetByEmailEmployee)
	router.GET("/get/all", leaveRequest.GetAll)
	router.POST("/create", leaveRequest.CreateOne)
	router.PUT("/update", leaveRequest.UpdateOne)
	router.PUT("/update/remaining", leaveRequest.UpdateRemainingLeaveDays)
	router.DELETE("/delete", leaveRequest.DeleteOne)
}

package leave_request_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	leaverequestcontroller "shop_erp_mono/api/controllers/human_resources_management/leave_request"
	"shop_erp_mono/bootstrap"
	employees_domain "shop_erp_mono/domain/human_resource_management/employees"
	leaverequestdomain "shop_erp_mono/domain/human_resource_management/leave_request"
	employeerepository "shop_erp_mono/repository/human_resource_management/employee/repository"
	leaverequestrepository "shop_erp_mono/repository/human_resource_management/leave_request/repository"
	leave_request_usecase "shop_erp_mono/usecase/human_resource_management/leave_request/usecase"
	"time"
)

func LeaveRequestRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup, cacheTTL time.Duration) {
	lr := leaverequestrepository.NewLeaveRequestRepository(db, leaverequestdomain.CollectionLeaveRequest)
	em := employeerepository.NewEmployeeRepository(db, employees_domain.CollectionEmployee)
	leaveRequest := &leaverequestcontroller.LeaveRequestController{
		LeaveRequestUseCase: leave_request_usecase.NewLeaveRequestUseCase(timeout, lr, em, cacheTTL),
		Database:            env,
	}

	router := group.Group("/leave-requests")
	router.GET("/get/_id", leaveRequest.GetByID)
	router.GET("/get/email", leaveRequest.GetByEmailEmployee)
	router.GET("/get/all", leaveRequest.GetAll)
	router.POST("/create", leaveRequest.CreateOne)
	router.PUT("/update", leaveRequest.UpdateOne)
	router.DELETE("/delete", leaveRequest.DeleteOne)
}

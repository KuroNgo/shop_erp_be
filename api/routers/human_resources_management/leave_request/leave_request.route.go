package leave_request_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	leaverequestcontroller "shop_erp_mono/api/controllers/human_resources_management/leave_request"
	"shop_erp_mono/bootstrap"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	leaverequestdomain "shop_erp_mono/domain/human_resource_management/leave_request"
	leaverequestrepository "shop_erp_mono/repository/human_resource_management/leave_request/repository"
	leave_request_usecase "shop_erp_mono/usecase/human_resource_management/leave_request/usecase"
	"time"
)

func LeaveRequestRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	lr := leaverequestrepository.NewLeaveRequestRepository(db, leaverequestdomain.CollectionLeaveRequest, employeesdomain.CollectionEmployee)
	leaveRequest := &leaverequestcontroller.LeaveRequestController{
		LeaveRequestUseCase: leave_request_usecase.NewLeaveRequestUseCase(timeout, lr),
		Database:            env,
	}

	router := group.Group("/leave_requests")
	router.GET("/get/_id", leaveRequest.FetchOneByIDLeaveRequest)
	router.GET("/get/email", leaveRequest.FetchOneByEmailLeaveRequest)
	router.GET("/get/all", leaveRequest.FetchAllLeaveRequest)
	router.POST("/create", leaveRequest.CreateOneLeaveRequest)
	router.PUT("/update", leaveRequest.UpdateOneLeaveRequest)
	router.DELETE("/delete", leaveRequest.DeleteOneLeaveRequest)
}